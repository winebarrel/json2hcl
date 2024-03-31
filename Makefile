.PHONY: all
all: vet test build

.PHONY: build
build:
	go build ./cmd/json2hcl

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -f json2hcl json2hcl.exe
