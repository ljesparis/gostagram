GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/gostagram_webapp_example

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) main.go

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web
