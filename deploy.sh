#!/usr/bin/env bash

#go fmt ./...
#git config credential.helper store
#git add . && git commit -am 'hello-0.0.1' && git push
#docker build --tag hello-0.0.1 .
#rm -rf vendor
#docker push hello-0.0.1

make tests
make lint
#make build2
docker build -t hello-0.0.7 .
docker-compose up -d
