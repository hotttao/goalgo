GOPATH:=$(shell go env GOPATH)
GOROOT:=$(shell go env GOROOT)
API_PROTO_FILES=$(shell find protobuf -name *.proto)


.PHONY: golang
# generate api proto
golang:
	protoc --proto_path=./protobuf \
	       --proto_path=../third_party \
 	       --go_out=paths=source_relative:./golang/api \
 	       --go-http_out=paths=source_relative:./golang/api \
 	       --go-grpc_out=paths=source_relative:./golang/api \
 	       --openapi_out==paths=source_relative:./golang \
	       $(API_PROTO_FILES)

.PHONY: python
# generate api proto
python:
	python -m grpc_tools.protoc \
		   --proto_path=./protobuf \
	       --proto_path=../third_party \
 	       --python_out=./python/ \
 	       --grpc_python_out=./python/ \
 	       --openapi_out==paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: all
# generate all
all:
	make golang;
	make python;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
