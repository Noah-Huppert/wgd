SRC_NAME = rpc
SRC_FILE = ${SRC_NAME}.proto

GO_OUT_DIR = .
GO_PB_OUT = ${GO_OUT_DIR}/${SRC_NAME}.pb.go

${GO_PB_OUT}: ${SRC_FILE}
	protoc --go_out=plugins=grpc:${GO_OUT_DIR} ${SRC_FILE}