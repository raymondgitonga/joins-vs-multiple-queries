.PHONY: build


default: build

make-down:
	docker-compose down

make-up:
	docker-compose up -d

tests:
	go test -v ./... | { grep -v 'no test files'; true; }

run:
	go run ./cmd

ci_lint:
	golangci-lint run ./... --fix

format:
	gofmt -w -s .

linter: format ci_lint

restart: make-down make-up run

build: make-up run