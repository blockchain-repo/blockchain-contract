#!/bin/bash
# The set -e option instructs bash to immediately exit
# if any command has a non-zero exit status
set -e

CUR_PATH=$(cd "$(dirname "$0")"; pwd)

protos_file_path=${CUR_PATH}/src/core/protos/

export PATH=/opt/protoc/bin:$PATH

proto_count=`find ${protos_file_path} -name *.proto|wc -l`

if [ "${proto_count}" != "0" ] ; then
    echo -e "\033[01;34m reproduce the *.go files according the *.proto files \033[1m"
    protoc -I=src/core/protos --go_out=src/core/protos/ src/core/protos/*.proto
    echo -e "\033[01;34m reproduce success! \033[1m"
    exit 0
else
    echo -e "\033[01;31m not exist the *.proto files in path ${protos_file_path} \033[01m"
fi



