#!/usr/bin/env bash

make vendor
make tests
make lint
git add . && git commit -am 'hello-0.0.9' && git push
