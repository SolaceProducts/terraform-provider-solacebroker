#!/usr/bin/env bash
set -euo pipefail

certs=$(cd "${1:?usage: run-cases.sh <cert-dir>}" && pwd)
here=$(cd "$(dirname "$0")" && pwd)
empty=$(mktemp -d)

# name | cert | provider GODEBUG | extra env | expected: success, or an error substring
cases=(
  "good|good|fips140=only|-|success"
  "rsa1024|rsa1024|fips140=only|-|no FIPS compatible certificate chains found"
  "rsa1024-control|rsa1024|fips140=off|-|success"
  "rsa1024-default|rsa1024|default|-|no FIPS compatible certificate chains found"
  "wrongname|wrongname|fips140=only|-|certificate is valid for wrong.example, not localhost"
  "cnonly|cnonly|fips140=only|-|relies on legacy Common Name field"
  "chain|chain|fips140=only|-|success"
  "leafonly|leafonly|fips140=only|-|certificate signed by unknown authority"
  "skip-verify|good|fips140=only|SOLACEBROKER_INSECURE_SKIP_VERIFY=true|cannot be true in FIPS mode"
)

serve_cert_and_wait() {
  local resp want got
  resp=$(jq -n --rawfile c "$certs/$1.pem" '{tlsServerCertContent: $c}' |
    curl -sS --fail-with-body -u admin:admin -X PATCH -H 'Content-Type: application/json' -d @- localhost:8080/SEMP/v2/config 2>&1) ||
    { echo "$resp" >&2; return 1; }
  want=$(openssl x509 -in "$certs/$1.crt" -noout -fingerprint -sha256) || return 1
  for _ in $(seq 60); do
    got=$(openssl s_client -connect localhost:1943 -servername localhost < /dev/null 2> /dev/null | openssl x509 -noout -fingerprint -sha256 2> /dev/null || true)
    [ "$got" = "$want" ] && return 0
    sleep 1
  done
  echo "broker never served the $1 cert" >&2
  return 1
}

tofu_trusting_test_root() {
  env SSL_CERT_FILE="$certs/root.crt" SSL_CERT_DIR="$empty" PROVIDER_GODEBUG="$godebug" ${extra:+"$extra"} \
    tofu -chdir="$here" "$@" -auto-approve -no-color -input=false
}

failed=0
for row in "${cases[@]}"; do
  IFS='|' read -r name cert godebug extra expected <<< "$row"
  [ "$extra" = - ] && extra=
  if ! serve_cert_and_wait "$cert"; then echo "FAIL $name (could not serve the $cert certificate)"; failed=1; continue; fi
  rc=0
  out=$(tofu_trusting_test_root apply 2>&1) || rc=$?
  if [ "$expected" = success ]; then
    ok=$([ "$rc" -eq 0 ] && echo 1 || echo 0)
    [ "$ok" = 1 ] && { out=$(tofu_trusting_test_root destroy 2>&1) || { rc=$?; ok=0; }; }
  else
    ok=$([ "$rc" -ne 0 ] && grep -qF "$expected" <<< "$(tr -s '[:space:]' ' ' <<< "$out")" && echo 1 || echo 0)
  fi
  if [ "$ok" = 1 ]; then echo "PASS $name"; else echo "FAIL $name (exit $rc, expected: $expected)"; echo "$out"; failed=1; fi
done
exit "$failed"
