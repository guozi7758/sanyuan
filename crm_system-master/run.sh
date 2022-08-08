#!/bin/bash

build(){
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o crm_$1 main.go
}

if [[ "" = $1 ]];then
        echo "端口号不可以为空"
	# shellcheck disable=SC2242
	exit -1
fi

build $1