LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml