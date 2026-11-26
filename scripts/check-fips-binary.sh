#!/usr/bin/env bash
# Checks a FIPS-capable provider binary. Run by the fips build's post hook in .goreleaser.yml.
set -euo pipefail
bin=$(realpath "$1")
cd "$(dirname "$0")/.."
# Repeated from .goreleaser.yml on purpose: a change there must also change this check.
pin=v1.0.0-c2097c7c

info=$(go version -m "$bin")
grep -qx $'\tbuild\tGOFIPS140='"$pin" <<< "$info" || { echo "$bin was not built with GOFIPS140=$pin" >&2; exit 1; }
grep -qx $'\tbuild\tCGO_ENABLED=0' <<< "$info" || { echo "$bin was not built with CGO_ENABLED=0" >&2; exit 1; }

# Crypto outside the standard library bypasses the validated module. By path only; other names are left to review.
deps=$(GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go list -deps -f '{{if not .Standard}}{{.ImportPath}}{{end}}' .)
crypto=$(grep crypto <<< "$deps" || true)
[ -z "$crypto" ] || { echo "the FIPS build uses crypto outside the standard library: ${crypto//$'\n'/ }" >&2; exit 1; }

# Fails on any known vulnerability in Go or a module the binary uses.
go run golang.org/x/vuln/cmd/govulncheck@v1.10.0 -mode=binary "$bin"
