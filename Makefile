# Makefile for dateparse project
.PHONY: codecov show test

CODE_COV_BASH := codecov.sh
CODE_COVERAGE := coverage.txt

codecov: test
	curl -o $(CODE_COV_BASH) -s https://codecov.io/bash
	chmod +x $(CODE_COV_BASH)
	./$(CODE_COV_BASH) -t b96414a5-3abe-4185-a8ce-6b482cea2753
	# rm -rf $(CODE_COV_BASH)

demo:
	go run example/main.go

dep:
	go get -t -v ./...

show: test
	go tool cover -html="$(CODE_COVERAGE)"

test:
	go test -race -coverprofile=$(CODE_COVERAGE) -covermode=atomic
