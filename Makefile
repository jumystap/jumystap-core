build:
	@go build -o bin/jumystap cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/jumystap
