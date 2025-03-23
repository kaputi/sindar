BINARY_NAME=sindar

check-quality:
	# make lint
	make fmt
	make vet

lint:
	golangci-lint run 

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

build:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux main.go
	@echo "Build complete"

run: build
	./bin/${BINARY_NAME}-linux

clean:
	go clean
	rm -rf bin

test-local:
	make tidy
	gotest -v ./...

test:
	make tidy
	gotest -v ./... -coverprofile=coverage.out -json > report.json

coverage:
	make test
	go tool cover -html=coverage.out

.PHONY: all test and build
all:
	make check-quality
	make test
	make build
