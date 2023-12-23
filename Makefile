build:
	@go build -o bin/theexpanse

run: build
	@./bin/theexpanse

test:
	@go test -v ./...
