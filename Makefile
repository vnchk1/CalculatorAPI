APP_NAME := calculator-api
MAIN := cmd/api/main.go

.PHONY: all build run swag fmt lint test clean tidy

all: build

build:
	@echo "Running build..."
	go build -o $(APP_NAME) $(MAIN)

run:
	@echo "Running the app..."
	go run $(MAIN)

swag:
	swag init -g $(MAIN) --output docs

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./... -v

clean:
	go clean
	if exist $(APP_NAME) del $(APP_NAME)
tidy:
	go mod tidy