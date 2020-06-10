BUILD_NAME=app

all: modules test build
modules:
	go mod download
test:
	go test -v ./...
build:
	go build -o $(BUILD_NAME) -v
build_linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_NAME) -v
run:
	go run .
validate_cci:
	circleci config validate .circleci/config.yml
