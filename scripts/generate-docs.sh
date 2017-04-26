#!/bin/bash
ROOT=$(dirname "${BASH_SOURCE}")/..

cd ${ROOT}
go run ./scripts/generate-docs.go
cd - > /dev/null
