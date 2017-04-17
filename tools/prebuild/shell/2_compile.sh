#!/bin/bash

# --------------------------------------------------
current_path=`pwd`
target="./bin/unicontract"
protos_file_path=$current_path/src/core/protos
# --------------------------------------------------

# --------------------------------------------------
#get govendor
echo "get govendor"
cd $current_path
go get github.com/kardianos/govendor
govendor sync
# --------------------------------------------------

# --------------------------------------------------
#proto-gen-go
echo "proto-gen-go"
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
proto_count=`find $protos_file_path -name *.proto | wc -l`

if [ $proto_count -ne 0 ]
then
    echo -e "reproduce the *.go files according the *.proto files"
    protoc -I=$protos_file_path --go_out=$protos_file_path $protos_file_path/*.proto
    echo -e "reproduce success!"
else
    echo -e "not exist the *.proto files in path $protos_file_path"
fi
# --------------------------------------------------

# --------------------------------------------------
#build
echo "build"
if [ $@ = "debug" ]
then
    go build -o $target && echo "ok"
else
    go build -ldflags "-s -w" -o $target && echo "ok"
fi
# --------------------------------------------------
