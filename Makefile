.PHONY: build


default: build

make-down:
	docker-compose down

make-up:
	docker-compose up -d

tests:
	go test -bench=. -v ./... | { grep -v 'no test files'; true; }

ci_lint:
	golangci-lint run ./... --fix

format:
	gofmt -w -s .

linter: format ci_lint

restart: make-down make-up

benchmark: make-up tests