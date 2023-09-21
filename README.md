# Test App

Test application

## Getting started

For test and deploy in local just use docker compose up -d orrun the program with go run main.go

## Create swagger documentation
You need to install swag and export the variable with this commands:
```sh
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$(go env GOPATH)/bin:$PATH
``` 

Run this to generate the documentation:
```sh
swag init -g main.go -o ./docs
```

## Test

To do the test of the project type
```sh
go test ./... -cover
```