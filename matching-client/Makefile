.PHONY: build


default: build

tests:
	go test -v ./... | { grep -v 'no test files'; true; }

run:
	go run ./cmd/web

ci_lint:
	golangci-lint run ./... --fix

format:
	gofmt -w -s .

linter: format ci_lint

build: run