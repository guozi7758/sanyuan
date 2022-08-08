#!/bin/bash

build_scp(){
	scp crm_$1 shane@106.13.75.157:/home/shane/server/$1
}

if [[ "" = $1 ]];then
	echo "端口号不可以为空"
	exit -1
fi

build_scp $1
