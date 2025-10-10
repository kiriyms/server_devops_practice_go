run: build
	@./bin/app

build:
	@go build -o bin/app ./cmd

devops:
	@act push --secret-file .secrets