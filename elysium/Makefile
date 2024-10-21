PROTOC_LINUX_VERSION = 3.11.4
PROTOC_LINUX_ZIP = protoc-$(PROTOC_LINUX_VERSION)-linux-x86_64.zip

BUF_VERSION=1.6.0
BUF_BINARY_NAME=buf

.PHONY: install-protoc install-protoc-go

install-protoc:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_LINUX_VERSION)/$(PROTOC_LINUX_ZIP)
	sudo unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_LINUX_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_LINUX_ZIP)

install-protoc-go:
	go install github.com/golang/protobuf/protoc-gen-go@v1.4.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.1
	go install github.com/envoyproxy/protoc-gen-validate@v0.4.1
	go install github.com/gogo/protobuf/protoc-gen-gofast@v1.3.1
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.7

install-buf:
	sudo curl -sSL "https://github.com/bufbuild/buf/releases/download/v$(BUF_VERSION)/$(BUF_BINARY_NAME)-$(shell uname -s)-$(shell uname -m)"  -o "/usr/local/bin/$(BUF_BINARY_NAME)" && sudo chmod +x "/usr/local/bin/$(BUF_BINARY_NAME)"


.PHONY: start-infra stop-infra

start-infra:
	docker compose -f ./dockerfiles/infrastructure.yaml up -d

stop-infra:
	docker compose -f ./dockerfiles/infrastructure.yaml down