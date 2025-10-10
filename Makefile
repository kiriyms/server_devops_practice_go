run: build
	@./bin/app

build:
	@go build -o bin/app ./cmd

test:
	@go test -v ./...

devops:
	@act push --secret-file .secrets