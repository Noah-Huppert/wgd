.PHONY: all

SRC_NAME = rpc
SRC_FILE = ${SRC_NAME}.proto

GO_OUT_DIR = .
GO_PB_OUT = ${GO_OUT_DIR}/${SRC_NAME}.pb.go

SERVER_DIR = ../server/rpc
CLIENT_DIR = ../client/rpc

${GO_PB_OUT}: ${SRC_FILE}
	protoc --go_out=plugins=grpc:${GO_OUT_DIR} ${SRC_FILE}
	if [ ! -d ${SERVER_DIR} ]; then mkdir ${SERVER_DIR}; fi
	if [ ! -d ${CLIENT_DIR} ]; then mkdir ${CLIENT_DIR}; fi
	cp ${GO_PB_OUT} ${SERVER_DIR}/${SRC_NAME}.pb.go
	cp ${GO_PB_OUT} ${CLIENT_DIR}/${SRC_NAME}.pb.go
