#!/usr/bin/env bash
set -euo pipefail
cd "$1"
ec=(-newkey ec -pkeyopt ec_paramgen_curve:P-256)
san='subjectAltName=DNS:localhost,IP:127.0.0.1'
ca() { # name signer|self cn
  openssl req -new "${ec[@]}" -nodes -keyout "$1.key" -out "$1.csr" -subj "/CN=$3"
  printf 'basicConstraints=critical,CA:TRUE\nkeyUsage=critical,keyCertSign\n' > "$1.ext"
  if [ "$2" = self ]; then openssl x509 -req -in "$1.csr" -key "$1.key" -out "$1.crt" -days 2 -sha256 -extfile "$1.ext"
  else openssl x509 -req -in "$1.csr" -CA "$2.crt" -CAkey "$2.key" -CAcreateserial -out "$1.crt" -days 2 -sha256 -extfile "$1.ext"; fi
}
leaf() { # name signer cn ext keyargs...
  local n=$1 s=$2 cn=$3 ext=$4; shift 4
  openssl req -new "$@" -nodes -keyout "$n.key" -out "$n.csr" -subj "/CN=$cn"
  printf '%s\nextendedKeyUsage=serverAuth\n' "$ext" > "$n.ext"
  openssl x509 -req -in "$n.csr" -CA "$s.crt" -CAkey "$s.key" -CAcreateserial -out "$n.crt" -days 2 -sha256 -extfile "$n.ext"
  cat "$n.crt" "$n.key" > "$n.pem"
}
ca root self a2-test-root
ca int root a2-test-int
leaf good      root localhost     "$san"                          "${ec[@]}"
leaf rsa1024   root localhost     "$san"                          -newkey rsa:1024
leaf wrongname root wrong.example 'subjectAltName=DNS:wrong.example' "${ec[@]}"
leaf cnonly    root localhost     'basicConstraints=CA:FALSE'     "${ec[@]}"
leaf leafonly  int  localhost     "$san"                          "${ec[@]}"
