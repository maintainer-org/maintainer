#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..

cd ${ROOT}
go run ./scripts/generate-docs.go
cp references/README.md references/maintainer.md
cd - > /dev/null
