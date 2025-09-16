LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.4.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.yaml

fmt:
	$(LOCAL_BIN)/golangci-lint fmt ./... --config .golangci.yaml