#!/usr/bin/env bash
set -e

apt-get -y gpm

gpm install

go get github.com/nsqio/nsq/...

exit 0