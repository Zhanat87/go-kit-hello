#!/usr/bin/env bash

make vendor
make tests
make lint
git add . && git commit -am 'hello-0.1.0' && git push
