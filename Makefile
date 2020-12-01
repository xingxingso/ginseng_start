
SHELL := /bin/bash

SRC_DIR=proto
DST_DIR=proto

.PHONY: proto # 防止文件名proto冲突
#添加`@` 不打印该命令信息
proto:
	@protoc -I=${SRC_DIR} --go_out=paths=source_relative:${DST_DIR} ${SRC_DIR}/*.proto