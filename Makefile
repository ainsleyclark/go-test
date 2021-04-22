lint:
	golangci-lint run ./...

format:
	go fmt ./...

test:
	go clean -testcache && go test -race ./...

mock:
	rm -rf mocks && mockery --all --keeptree

all:
	$(MAKE) format
	$(MAKE) lint
	$(MAKE) test