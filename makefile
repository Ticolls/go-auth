.PHONY: default run build test docs clean

# variables
APP_NAME=go-auth 

# Tasks
default: run-with-docs

run:
	@/usr/local/go/bin/go run main.go

run-with-docs:
	@swag init
	@/usr/local/go/bin/go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./ ...

docs: 
	@swag init

clean: 
	@rm -f $(APP_NAME)
	@rm -rf ./docs