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
goos=$(awk '/^ *goos:/ {on=1; next} on && /^ *(#|$)/ {next} on && /^ *- / {print $2; next} on {exit}' .goreleaser.yml)
for os in linux windows; do
  grep -qx "$os" <<< "$goos" || { echo "$os is missing from the goos list read from .goreleaser.yml: ${goos//$'\n'/ }" >&2; exit 1; }
done

for os in $goos; do
  GOOS=$os GOARCH=amd64 go list -deps -f '{{with .Module}}{{if not .Main}}{{.Path}} {{.Version}} {{.Dir}}{{end}}{{end}}' . >> "$tmp/modules"
  GOOS=$os GOARCH=amd64 "$tmp/go-licenses" csv . >> "$tmp/licenses.csv" 2>> "$tmp/go-licenses.log" \
    || { cat "$tmp/go-licenses.log" >&2; exit 1; }
done
LC_ALL=C sort -u "$tmp/modules" -o "$tmp/modules"

# SPDX id, its heading under texts_heading, and the last line of its text there (so a cut-off text fails).
cat > "$tmp/licences" << 'EOF'
Apache-2.0|Apache 2.0|END OF TERMS AND CONDITIONS
BSD-2-Clause|BSD 2|SUCH DAMAGE.
BSD-3-Clause|BSD 3|SUCH DAMAGE.
MIT|MIT|OTHER DEALINGS IN THE SOFTWARE.
MPL-2.0|Mozilla 2.0|defined by the Mozilla Public License, v. 2.0.
EOF

sed -n "/^$texts_heading\$/,\$p" "$inventory" > "$tmp/texts"
[ -s "$tmp/texts" ] || { echo "$inventory has no '$texts_heading' section" >&2; exit 1; }

# go-licenses names a licence per package root (golang.org/x/sys/unix); give each module the licence of the
# roots inside it. Its link column needs the network and isn't used.
{
  echo 'THIRD-PARTY SOFTWARE USED'
  echo '========================='
  echo
  while read -r path version dir; do
    spdx=$(awk -F, -v m="$path" '$1 == m || index($1, m "/") == 1 {print $3}' "$tmp/licenses.csv" | LC_ALL=C sort -u)
    [ -n "$spdx" ] || { echo "go-licenses found no licence for $path" >&2; exit 1; }
    names=""
    for s in $spdx; do
      name=$(awk -F'|' -v s="$s" '$1 == s {print $2}' "$tmp/licences")
      [ -n "$name" ] || { echo "$path is $s: add it to the licence table in $0 and its text to $inventory" >&2; exit 1; }
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
    echo "${copyright:-No copyright line in LICENSE.}"
    echo "Home page: https://pkg.go.dev/$path@$version"
    echo
  done < "$tmp/modules"
} > "$tmp/modules.txt"
LC_ALL=C sort -u "$tmp/used" -o "$tmp/used"

# Each licence used needs its heading under texts_heading and, before the next heading, the last line of its text.
missing=$(awk -F'|' 'FILENAME == ARGV[1] {last[$2] = $3; next}
  FILENAME == ARGV[2] {if ($0 in last) cur = $0; else if (cur != "" && index($0, last[cur])) ok[cur] = 1; next}
  !($0 in ok)' "$tmp/licences" "$tmp/texts" "$tmp/used")
[ -z "$missing" ] || { echo "the text of $(paste -sd, - <<< "$missing") under '$texts_heading' in $inventory is missing or cut off" >&2; exit 1; }

{
  echo 'LICENSE SUMMARY'
  echo '==============='
  echo
  echo 'License terms can be found at the bottom of this file.'
  echo
  cat "$tmp/used"
  echo
  echo
  cat "$tmp/modules.txt"
  echo
  cat "$tmp/texts"
} > "$tmp/inventory"

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
} > "$tmp/notice"

mv "$tmp/inventory" "$inventory"
mv "$tmp/notice" "$notice"
