#!/bin/bash
# The set -e option instructs bash to immediately exit
# if any command has a non-zero exit status
set -e

CUR_PATH=$(cd "$(dirname "$0")"; pwd)
export GOPATH=${CUR_PATH}:$GOPATH

#bee version -o json

#go get -u github.com/kardianos/govendor
go get github.com/kardianos/govendor

#govendor init
govendor init
govendor add +external

#app run
bee run
