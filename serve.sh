#!/bin/bash

# 获取 .env 文件里配置的端口号
source .env
port=$APP_PORT

# 监听端口号是否被占用
if netstat -an | grep "$port" | grep -i listen > /dev/null ; then
    lsof -i:"$port" | grep main
    echo -e "\033[1;31m$port 端口被占用，请输入被占用的进程PID: \033[0m"
    read -r  pid
    kill -15 "$pid"
    if [ $? -eq 0 ]; then
        echo -e "\033[1;32m进程($pid)终止成功，正重启服务··· \033[0m"
        sleep 1
        air
    else
        echo -e "\033[1;31m进程($pid)终止失败，请检查输入的进程ID是否正确 \033[0m"
    fi
else
   air
fi