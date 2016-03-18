#!/bin/bash

set -e

COVER=cover
ROOT_PKG=github.com/thriftrw/thriftrw-go/

if [[ -d "$COVER" ]]; then
	rm -rf "$COVER"
fi
mkdir -p "$COVER"

i=0
for pkg in "$@"; do
	i=$((i + 1))

	coverpkg=$(go list -json "$pkg" | jq -r '
		.Deps
		| map(select(startswith("'"$ROOT_PKG"'")))
		| . + ["'"$pkg"'"]
		| join(",")
	')

	go test \
		-coverprofile "$COVER/cover.${i}.out" -coverpkg "$coverpkg" \
		-v "$pkg"
done

gocovmerge "$COVER"/*.out > cover.out