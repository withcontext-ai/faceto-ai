GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git | grep cmd))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: docker-build
# generate all
docker-build:
	docker build -f ./server/Dockerfile -t "faceto-ai-server:v$(VERSION)" .
	docker build -f ./web/Dockerfile -t "faceto-ai-web:v$(VERSION)" .
	docker images

.PHONY: docker-run
# generate all
docker-run:
	docker run -it -dp 8001:8001 "faceto-ai-server:v$(VERSION)"
	docker run -it -dp 3000:3000 "faceto-ai-web:v$(VERSION)"
	docker container ls

.PHONY: docker-image-rm
# generate all
docker-image-rm:
	docker image rm "faceto-ai-server:v$(VERSION)"
	docker image rm "faceto-ai-web:v$(VERSION)"

.PHONY: docker-all
# generate all
docker-all:
	make docker-build
	make docker-run

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
