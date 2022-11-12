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

build: docker-compose-up run