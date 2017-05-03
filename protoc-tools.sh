#!/usr/bin/env bash

current_path=$(cd `dirname $0`; pwd)
#默认的proto schema文件路径
protos_file_path=${current_path}/src/core/protos

# remove the string <,omitempty> in generated .pb.go files[proto3]
remove_omitempty=true

proto_support_array=("go","java","js","python")
#proto_support_array=("go","java","js")

proto_js_libs_name="proto_contract_libs"

if [[ $# -le 0 ]]; then
    echo -e "input args should in ${proto_support_array}"
fi

function generateProto ()
{
    if [ $# == 1 ]; then
        outputDir=${protos_file_path}
        if [ "$@" != "go" ];then
            outputDir=${protos_file_path}/"$@"
        fi

        if ! [ -d  ${outputDir} ]; then
            echo "${outputDir} not exist, will create it."
            mkdir -p ${outputDir}
        fi

        echo "proto-gen-$@"
        echo -e "reproduce the *.go files according to the ${protos_file_path}/*.proto files"
        if [ "$@" == "js" ]; then
#            protoc -I=${protos_file_path} --js_out=library=${proto_js_libs_name},binary:${outputDir} contract.proto

#            protoc --js_out=library=${proto_js_libs_name},binary:${outputDir} ${protos_file_path}/*.proto
            echo -e "protoc --js_out=library=${proto_js_libs_name},binary:${outputDir} ${protos_file_path}/*.proto"
            protoc -I=${protos_file_path} --js_out=library=${proto_js_libs_name},binary:${outputDir} ${protos_file_path}/*.proto
        else
            protoc -I=${protos_file_path} --"$@"_out=${outputDir} ${protos_file_path}/*.proto
        fi
        echo -e "reproduce success!"

        if [ "$@" == "go" ] && [ ${remove_omitempty} == true ]; then
            echo -e "replace the <,omitempty> in ${protos_file_path}/*.pb.go files"
            sed  -i 's/,omitempty//g' ${protos_file_path}/*.pb.go
            echo -e "replace the <,omitempty> success!"
        fi
    fi
}

validCount=0
for var in "$@"
do
    if [[ "${proto_support_array[@]}" =~  ${var} ]]; then
#        validCount=$[validCount=${validCount}+1]
        ((validCount++))
    fi
done

if [[ ${validCount} != "0" ]]; then
    go get github.com/golang/protobuf/{proto,protoc-gen-go}
    proto_count=`find $protos_file_path -name *.proto | wc -l`
    if [ $proto_count -ne 0 ]; then
        for var in "$@"
        do
            if [[ "${proto_support_array[@]}" =~  ${var} ]]; then
                generateProto ${var}
            fi
        done
	else
        echo -e "not exist the *.proto files in path $protos_file_path"
	fi
fi
exit 0



