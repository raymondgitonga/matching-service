.PHONY: build


default: build

docker-compose-down:
	docker-compose down

docker-compose-up:
	docker-compose up -d

tests:
	go test -v ./... | { grep -v 'no test files'; true; }

run:
	go run ./cmd/web

lint:
	golangci-lint run

build: docker-compose-up run