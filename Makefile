run: build
	@./bin/app

build:
	@go build -ldflags="-X main.commit=local" -o bin/app ./cmd

test:
	@go test -v ./...

devops.push:
	@act push --secret-file .secrets

devops.test:
	@act -j test

docker.build:
	@docker build --build-arg GIT_SHA=docker.local -t server_devops_practice_go:local .

docker.run:
	@docker run --rm -p 8080:8080 --env-file .prod.env server_devops_practice_go:local