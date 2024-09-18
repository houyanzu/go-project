#!/bin/bash

# 检查参数是否存在
if [ -z "$1" ] || [ -z "$2" ]; then
    echo "Usage: script.sh <arg1> <arg2>"
    exit 1
fi

# 参数解析
case $2 in
    "api")
        ./boxgen --action routergen
        apiOut="./bin/api.linux.$1"
        go build -o "$apiOut" ./app/api/home/main.go ./app/api/home/register.go
        ;;
    *)
        echo "wrong"
        ;;
esac
