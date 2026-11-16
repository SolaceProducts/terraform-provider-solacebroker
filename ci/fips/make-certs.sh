#!/usr/bin/env bash
set -euo pipefail
cd "${1:?usage: make-certs.sh <cert-dir>}"
ec=(-newkey ec -pkeyopt ec_paramgen_curve:P-256)
san='subjectAltName=DNS:localhost,IP:127.0.0.1'
ca() {
  local n=$1 signer=$2 cn=$3
  openssl req -new "${ec[@]}" -nodes -keyout "$n.key" -out "$n.csr" -subj "/CN=$cn"
  printf 'basicConstraints=critical,CA:TRUE\nkeyUsage=critical,keyCertSign\n' > "$n.ext"
  if [ "$signer" = self ]; then openssl x509 -req -in "$n.csr" -key "$n.key" -out "$n.crt" -days 2 -sha256 -extfile "$n.ext"
  else openssl x509 -req -in "$n.csr" -CA "$signer.crt" -CAkey "$signer.key" -CAcreateserial -out "$n.crt" -days 2 -sha256 -extfile "$n.ext"; fi
}
leaf() {
  local n=$1 signer=$2 cn=$3 ext=$4; shift 4
  openssl req -new "$@" -nodes -keyout "$n.key" -out "$n.csr" -subj "/CN=$cn"
  printf '%s\nextendedKeyUsage=serverAuth\n' "$ext" > "$n.ext"
  openssl x509 -req -in "$n.csr" -CA "$signer.crt" -CAkey "$signer.key" -CAcreateserial -out "$n.crt" -days 2 -sha256 -extfile "$n.ext"
  cat "$n.crt" "$n.key" > "$n.pem"
}
ca root self a2-test-root
ca int root a2-test-int
leaf good      root localhost     "$san"                          "${ec[@]}"
leaf rsa1024   root localhost     "$san"                          -newkey rsa:1024
leaf wrongname root wrong.example 'subjectAltName=DNS:wrong.example' "${ec[@]}"
leaf cnonly    root localhost     'basicConstraints=CA:FALSE'     "${ec[@]}"
leaf leafonly  int  localhost     "$san"                          "${ec[@]}"
