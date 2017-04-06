#!/bin/bash
set -e
#下载protoc3
wget https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip -c -P /opt
mkdir /opt/protoc

#下载zip
apt-get install zip

#解压zip包
unzip /opt/protoc-3.0.0-linux-x86_64.zip -d /opt/protoc
chomd +x /opt/protoc/

#配置环境变量
echo export PATH=$PATH:/opt/protoc/bin >>/etc/profile

source /etc/profile

#The compiler plugin protoc-gen-go will be installed in $GOBIN, defaulting to $GOPATH/bin
go get -u github.com/golang/protobuf/protoc-gen-go

exit 0