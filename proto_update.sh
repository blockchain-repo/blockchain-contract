#!/usr/bin/env bash

# todo

repo_version='master'
repo_name='uniledger-protos'
repo_url='http://36.110.71.170:99/uniinterface/uniledger-protos.git'

default_proto_project_dir='~/home/uniinterface/'

cd ${default_proto_project_dir}

mkdir -p ${default_proto_project_dir}

is_exist_repo=false

cd ${repo_name}

repo_log_count=`git log -n 5 --oneline|wc -l`

if [[  repo_log_count gt 1 ]]; then
    repo_log_count=true
fi

git clone -b ${repo_version} ${repo_name}