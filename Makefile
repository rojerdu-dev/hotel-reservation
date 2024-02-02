BINARY_NAME=api

build:
	@GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}

run: build
	@./bin/api

clean:
	@go clean
	@rm bin/${BINARY_NAME}

test:
	@go test -v ./...

test_coverage:
	@go test ./... -coverprofile=coverage.out

vet:
	@go vet

lint:
	@golangci-lint run