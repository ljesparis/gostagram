.PHONY: install reinstall clean test

install: # installing dependencies.
	@echo Installing mapstructure, gorequest, http2curl errors net text dependencies
	@go get github.com/mitchellh/mapstructure
	@go get github.com/parnurzeal/gorequest
	@go get github.com/moul/http2curl
	@go get github.com/pkg/errors
	@go get golang.org/x/net
	@go get golang.org/x/text

# reinstalling project from scratch.
reinstall: clean install

clean:
	@echo cleaning project.
	@rm -rf Gopkg.lock Gopkg.toml vendor

test:
	@go test -v -parallel 8 -tags gostagram
