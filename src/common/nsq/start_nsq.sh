#!/usr/bin/env bash

#管理nsqd信息
nsqlookupd &

#启动nsqd 负责接受,发送消息
nsqd --lookupd-tcp-address=127.0.0.1:4160 &

#启动web信息监控页面
nsqadmin --lookupd-http-address=127.0.0.1:4161 &