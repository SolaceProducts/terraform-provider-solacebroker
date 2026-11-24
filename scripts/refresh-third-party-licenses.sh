#!/usr/bin/env bash
# Regenerates third-party-licenses.txt and NOTICE from the modules built into the provider.
# The release runs this and fails if either file changes, so run it (make refresh-third-party-licenses)
# and commit the result whenever go.mod changes.
set -euo pipefail
cd "$(dirname "$0")/.."

inventory=third-party-licenses.txt
notice=NOTICE
texts_heading='LICENSE REQUIREMENTS & SPECIFICATIONS'

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

# Built for the host, so GOOS below only changes what it analyses.
GOBIN="$tmp" go install github.com/google/go-licenses@v1.6.0

# shortcut: the module set depends on GOOS, not GOARCH (checked 2026-11-24), so amd64 stands in for every
# architecture. Recheck if .goreleaser.yml gains a platform.
goos=$(awk '/^ *goos:/ {on=1; next} on && /^ *- / {print $2; next} on {exit}' .goreleaser.yml)
[ -n "$goos" ] || { echo "no goos list found in .goreleaser.yml" >&2; exit 1; }

for os in $goos; do
  GOOS=$os GOARCH=amd64 go list -deps -f '{{with .Module}}{{if not .Main}}{{.Path}} {{.Version}} {{.Dir}}{{end}}{{end}}' . >> "$tmp/modules"
  GOOS=$os GOARCH=amd64 "$tmp/go-licenses" csv . >> "$tmp/licenses.csv" 2>> "$tmp/go-licenses.log" \
    || { cat "$tmp/go-licenses.log" >&2; exit 1; }
done
LC_ALL=C sort -u "$tmp/modules" -o "$tmp/modules"

# go-licenses names a licence per package root (golang.org/x/sys/unix); give each module the licence of the
# roots inside it. Its link column needs the network and isn't used.
heading() {
  case $1 in
    Apache-2.0) echo 'Apache 2.0' ;;
    BSD-2-Clause) echo 'BSD 2' ;;
    BSD-3-Clause) echo 'BSD 3' ;;
    MIT) echo 'MIT' ;;
    MPL-2.0) echo 'Mozilla 2.0' ;;
    *) return 1 ;;
  esac
}

sed -n "/^$texts_heading\$/,\$p" "$inventory" > "$tmp/texts"
[ -s "$tmp/texts" ] || { echo "$inventory has no '$texts_heading' section" >&2; exit 1; }

{
  echo 'THIRD-PARTY SOFTWARE USED'
  echo '========================='
  echo
  while read -r path version dir; do
    spdx=$(awk -F, -v m="$path" '$1 == m || index($1, m "/") == 1 {print $3}' "$tmp/licenses.csv" | LC_ALL=C sort -u)
    [ -n "$spdx" ] || { echo "go-licenses found no licence for $path" >&2; exit 1; }
    names=""
    for s in $spdx; do
      name=$(heading "$s") || { echo "$path is $s: add it to heading() and its text to $inventory" >&2; exit 1; }
      grep -qxF "$name" "$tmp/texts" || { echo "$path is $name: add its text under '$texts_heading'" >&2; exit 1; }
      echo "$name" >> "$tmp/used"
      names="${names:+$names, }$name"
    done
    title="$path-$version"
    echo "$title"
    printf '%*s\n' "${#title}" '' | tr ' ' -
    echo "Licensed under $names"
    copyright=$(cat "$dir"/LICENSE* "$dir"/COPYING* 2> /dev/null \
      | grep -E '^[[:space:]]*Copyright ' | grep -vE '\[yyyy\]|<year>|\{yyyy\}' \
      | sed -E 's/^[[:space:]]+//; s/[[:space:]]+$//' | awk '!seen[$0]++' || true)
    echo "${copyright:-No copyright.}"
    echo "Home page: https://pkg.go.dev/$path@$version"
    echo
  done < "$tmp/modules"
} > "$tmp/modules.txt"

{
  echo 'LICENSE SUMMARY'
  echo '==============='
  echo
  echo 'License terms can be found at the bottom of this file.'
  echo
  LC_ALL=C sort -u "$tmp/used"
  echo
  echo
  cat "$tmp/modules.txt"
  echo
  cat "$tmp/texts"
} > "$tmp/inventory"
mv "$tmp/inventory" "$inventory"

{
  cat << 'EOF'
Terraform Provider for Solace Software Event Broker
Copyright 2024 Solace Corporation

Licensed under the Apache License, Version 2.0. See the LICENSE file for the
full terms.

--------------------------------------------------------------------------------
Third-party software

This product includes third-party software. Each component, its version, and
its license are listed in third-party-licenses.txt.

--------------------------------------------------------------------------------
Required attributions propagated from bundled dependencies

Apache License 2.0 section 4(d) requires that we reproduce the attribution
notices carried in a dependency's own NOTICE file. The blocks below are those
notices as upstream states them.
EOF
  while read -r path _ dir; do
    for f in "$dir"/NOTICE "$dir"/NOTICE.txt "$dir"/NOTICE.md; do
      [ -f "$f" ] || continue
      echo
      echo "$path"
      sed -E 's/^(.)/  \1/; s/[[:space:]]+$//' "$f"
    done
  done < "$tmp/modules"
} > "$notice"
