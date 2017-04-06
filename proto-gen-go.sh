#!/bin/bash
# The set -e option instructs bash to immediately exit
# if any command has a non-zero exit status
set -e

if [[ $# -eq 1 && $1 == "init" ]];then
    echo -e "\033[42;37m reproduce the *.go files according the *.proto files \033[0m"
    protoc --proto_path=src/core/protos --go_out=src/core/model/ src/core/protos/*.proto
    echo -e "\033[42;37m reproduce over, please rename the *.go files package \033[0m"
    exit 0
else
    echo -e  "\033[31m If you want reproduce the *.go files, please input param init!\033[0m"
    exit 1
fi


