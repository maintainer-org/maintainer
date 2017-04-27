#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..

cd ${ROOT}

./scripts/build-for-alpine.sh
docker build -t gaocegege/maintainer .
docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
docker push gaocegege/maintainer

cd - > /dev/null
