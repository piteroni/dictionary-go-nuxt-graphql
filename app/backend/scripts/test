#!/usr/bin/env sh

set -eu

cd "$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"

[ ! -d tmp ] && mkdir tmp
gotest -v -cover ./... -coverprofile=tmp/cover.out
go tool cover -html=tmp/cover.out -o tmp/cover.html
