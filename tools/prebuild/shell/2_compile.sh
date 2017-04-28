#!/bin/bash

# --------------------------------------------------
current_path=`pwd`
#默认的编译输出路径及目标文件名称
target="./bin/unicontract"
#默认的proto schema文件路径
protos_file_path=$current_path/src/core/protos

# remove the string <,omitempty> in generated .pb.go files[proto3]
remove_omitempty=true

# --------------------------------------------------

# --------------------------------------------------
#proto-gen-go
if [ $@ = "protoc" ]
then
	echo "proto-gen-go"
	go get github.com/golang/protobuf/{proto,protoc-gen-go}
	proto_count=`find $protos_file_path -name *.proto | wc -l`

	if [ $proto_count -ne 0 ]
	then
        echo -e "reproduce the *.go files according to the ${protos_file_path}/*.proto files"
        protoc -I=$protos_file_path --go_out=$protos_file_path $protos_file_path/*.proto
        echo -e "reproduce success!"
        if [ ${remove_omitempty} == true ]; then
            echo -e "replace the <,omitempty> in ${protos_file_path}/*.pb.go files"
            sed  -i 's/,omitempty//g' ${protos_file_path}/*.pb.go
            echo -e "replace the <,omitempty> success!"
        fi
	else
        echo -e "not exist the *.proto files in path $protos_file_path"
	fi
	exit 0
fi
# --------------------------------------------------

# --------------------------------------------------
#get govendor
echo "get govendor"
cd $current_path
go get github.com/kardianos/govendor
# 因为已经将三方包放在git上了，所以暂时不需要每次都进行sync
#govendor sync
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
