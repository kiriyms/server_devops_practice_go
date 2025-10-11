run: build
	@./bin/app

build:
	@go build -o bin/app ./cmd

test:
	@go test -v ./...

devops.push:
	@act push --secret-file .secrets

devops.test:
	@act -j test