GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto
proto:
	@protoc *.proto --go_out=. --go-grpc_out=.

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t loginsvc:latest .

.PHONY: check
check:
	@goctl env check --install --verbose --force