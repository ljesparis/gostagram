DEPPATH = $(shell which dep || echo 1)
GOLANGPACKAGE = $(shell test -s $(pwd)/Gopkg.lock && test -s $(pwd)/Gopkg.toml && echo 1)

install:
# checking if dep package manager(https://github.com/golang/dep)
# exists in the systems. If it doesn't exists, will be install it.
ifeq ($(DEPPATH),1)
	@echo installing golang package manager.
	@go get -u github.com/golang/dep
else
	@echo echo golang package mananger already installed.
endif

# checking if go package was initialized, if not, let create a go
# package.
ifeq ($(GOLANGPACKAGE),1)
	@echo initializing golang package.
	@dep init
else
	@echo golang package already initialized.
endif

	@echo installing dependencies.
	@dep ensure
	@echo installation finished.

.PHONY: install

reinstall:
	@make clean
	@make install

.PHONY: reinstall

clear:
#deleting Gopkg.* files, if they exists.
ifeq ($(GOLANGPACKAGE),1)
	@echo deleting Gopkg.* files
	@rm Gopkg.*
endif

# deleting vendor package.
ifeq ($(test -s $(pwd)/vendor && echo 1),1)
	@echo deleting vendor files.
	@rm -rf vendor
endif

.PHONY: clear

test:
	@echo test

.PHONY: test
