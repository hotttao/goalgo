GOPATH:=$(shell go env GOPATH)
GOROOT:=$(shell go env GOROOT)
API_PROTO_FILES=$(shell find api -name *.proto)


.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=../third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
 	       --openapi_out==paths=source_relative:. \
	       $(API_PROTO_FILES)


.PHONY: all
# generate all
all:
	make api;

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
