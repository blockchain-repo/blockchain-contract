#!/bin/bash
export GOPATH=/home/wangxin/company/golang/src/unicontract:$GOPATH

#bee version -o json

#govendor init
govendor init
govendor add +external

#app run
bee run
