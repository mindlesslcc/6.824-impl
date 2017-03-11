
all: build

build:proto

proto:
	protoc --proto_path=src/proto/ -I src/proto/ src/proto/*.proto --go_out=plugins=grpc:src/proto

test: build
	GOPATH="$(CURDIR)" && go test ./src/mapreduce
