#!/bin/bash
# The set -e option instructs bash to immediately exit
# if any command has a non-zero exit status
set -e

protos_file_path=src/core/protos/

export PATH=/opt/protoc/bin:$PATH

if [ -f ${protos_file_path}/*.proto ]; then
    echo -e "\033[01;34m reproduce the *.go files according the *.proto files \033[1m"
    protoc -I=src/core/protos --go_out=src/core/protos/ src/core/protos/*.proto
    echo -e "\033[01;34m reproduce success! \033[1m"
    exit 0
else
    echo -e "\033[01;31m not exist the *.proto files in path ${protos_file_path} \033[01m"
fi



