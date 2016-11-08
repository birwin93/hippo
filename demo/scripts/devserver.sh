#!/bin/bash

HOME_PATH="$GOPATH/src/github.com/birwin93/hippo/demo"

go run $HOME_PATH/*.go \
 -env=development \
 -config=$HOME_PATH/config.yml
