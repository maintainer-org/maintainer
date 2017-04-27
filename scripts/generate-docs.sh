#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..

cd ${ROOT}
go run ./scripts/generate-docs.go
cp docs/README.md docs/maintainer.md
cd - > /dev/null
