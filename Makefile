GOCMD=go
DOCKERCMD=docker
DOCKERCOMPOSECMD=docker-compose
SSL=openssl genrsa
GO_RUN=$(GOCMD) run
GO_BUILD=$(GOCMD) build
DOCKER_BUILD=$(DOCKERCMD) build
DOCKER_RUN=$(DOCKERCMD) run

GQLGEN=github.com/99designs/gqlgen

run:
	$(GO_RUN) ./cmd/main.go
compose-up:
	$(DOCKERCOMPOSECMD) -f ./build/docker-compose.yml up -d
compose-down:
	$(DOCKERCOMPOSECMD) -f ./build/docker-compose.yml down
docker-build:
	$(DOCKER_BUILD) ./ -t miraikeitai2020/bff:0.3.0
docker-run:
	$(DOCKER_RUN) -d -p 9020:9020 miraikeitai2020/bff:0.3.0