.PHONY: build build-alpine clean test help default list ci push

SHELL=/bin/bash

username             := bossjones
container_name       := go-chatbot-lab
BIN_NAME             := go-chatbot-lab
SOURCES              := $(shell find . \( -name vendor \) -prune -o  -name '*.go')
# VERSION              := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
VERSION              := $(shell git describe --tags --dirty)
GIT_SHA              := $(shell git rev-parse HEAD)
GIT_DIRTY            := $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_BRANCH           := $(shell git rev-parse --abbrev-ref HEAD)
IMAGE_NAME           := $(username)/$(container_name)
SOURCES              := $(shell find . \( -name vendor \) -prune -o  -name '*.go')
# LOCAL_REPOSITORY = $(HOST_IP):5000


mkfile_path          := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir          := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

define ASCICHATBOT
============GO CHATBOT LAB============
endef

export ASCICHATBOT

# http://misc.flogisoft.com/bash/tip_colors_and_formatting

RED=\033[0;31m
GREEN=\033[0;32m
ORNG=\033[38;5;214m
BLUE=\033[38;5;81m
PURP=\033[38;5;129m
GRAY=\033[38;5;246m
NC=\033[0m

export RED
export GREEN
export NC
export ORNG
export BLUE
export PURP
export GRAY

TAG ?= $(VERSION)
ifeq ($(TAG),@branch)
	override TAG = $(shell git symbolic-ref --short HEAD)
	@echo $(value TAG)
endif

# verify that certain variables have been defined off the bat
check_defined = \
    $(foreach 1,$1,$(__check_defined))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $(value 2), ($(strip $2)))))

list_allowed_args := name inventory

export PATH := ./bin:./venv/bin:$(PATH)

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

default: help

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
# re: ldflags - Link, typically invoked as “go tool link,” reads the Go archive or object for a package main, along with its dependencies, and combines them into an executable binary.
# SOURCE: https://golang.org/cmd/link/
	go build -ldflags "-X github.com/bossjones/go-chatbot-lab/shared/version.GitCommit=$$(git rev-parse HEAD) \
			          -X github.com/bossjones/go-chatbot-lab/shared/version.Version=$$(git describe --tags --dirty) \
			          -X github.com/bossjones/go-chatbot-lab/shared/version.VersionPrerelease=DEV \
			          -X github.com/bossjones/go-chatbot-lab/shared/version.BuildDate=$$(date -u +'%FT%T%z')" \
	-o bin/${BIN_NAME}

#REQUIRED-CI
bin/go-chatbot-lab: $(SOURCES)
	@if [ "$$(uname)" == "Linux" ]; then \
		CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -i -v \
			-ldflags "-X github.com/bossjones/go-chatbot-lab/shared/version.GitCommit=$$(git rev-parse HEAD) \
			          -X github.com/bossjones/go-chatbot-lab/shared/version.Version=$$(git describe --tags --dirty) \
			          -X github.com/bossjones/go-chatbot-lab/shared/version.VersionPrerelease=DEV \
			          -X github.com/bossjones/go-chatbot-lab/shared/version.BuildDate=$$(date -u +'%FT%T%z')" \
			-o bin/go-chatbot-lab \
			./ ;  \
	else \
		echo "Skipping Linux CGO build, because \"uname\" returned \"$$(uname)\"" ; \
	fi

install-deps: install-tools get-deps

get-deps:
	glide install

update:
	glide update

glide-install-update: get-deps update

install-tools:
	@which golint || go get -u github.com/golang/lint/golint
	@which cover || go get golang.org/x/tools/cmd/cover
	@test -d $$GOPATH/github.com/go-ini/ini || go get github.com/go-ini/ini
	@test -d $$GOPATH/github.com/jmespath/go-jmespath ||  go get github.com/jmespath/go-jmespath
	@which ginkgo || go get github.com/onsi/ginkgo/ginkgo
	@which gomega || go get github.com/onsi/gomega
	@which gomock || go get github.com/golang/mock/gomock
	@which mockgen || go get github.com/golang/mock/mockgen
	@which glide || go get github.com/Masterminds/glide
	@which go-bindata || go get -u github.com/jteeuwen/go-bindata/...
	@which gocov || go get github.com/axw/gocov/gocov
	@which gocov-xml || go get github.com/AlekSi/gocov-xml
	@which godepgraph || go get github.com/kisielk/godepgraph
	@which goveralls || go get github.com/mattn/goveralls
	@which gover || go get github.com/modocache/gover

build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X main.GitCommit=${GIT_SHA}${GIT_DIRTY} -X main.VersionPrerelease=VersionPrerelease=RC' -o bin/${BIN_NAME}

package:
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_SHA)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_SHA=$(GIT_SHA) -t $(IMAGE_NAME):local .

tag:
	@echo "Tagging: latest ${VERSION} $(GIT_SHA)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_SHA)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_SHA)"
	docker push $(IMAGE_NAME):$(GIT_SHA)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

# INFO: glide nv = List all non-vendor paths in a directory.
# test:
# 	go test $(glide nv)

#REQUIRED-CI
non_docker_compile:  install-deps bin/go-chatbot-lab

non_docker_lint: install-deps
	go tool vet -all config shared
	@DIRS="config/... shared/..." && FAILED="false" && \
	echo "gofmt -l *.go config shared" && \
	GOFMT=$$(gofmt -l *.go config shared) && \
	if [ ! -z "$$GOFMT" ]; then echo -e "\nThe following files did not pass a 'go fmt' check:\n$$GOFMT\n" && FAILED="true"; fi; \
	for codeDir in $$DIRS; do \
		echo "golint $$codeDir" && \
		LINT="$$(golint $$codeDir)" && \
		if [ ! -z "$$LINT" ]; then echo "$$LINT" && FAILED="true"; fi; \
	done && \
	if [ "$$FAILED" = "true" ]; then exit 1; else echo "ok" ;fi

#REQUIRED-CI
non_docker_ginko_cover:
	.ci/ginko-cover

#REQUIRED-CI
# FIXME: Skip this for now cause of non_docker_compile (12/6/2017)
# non_docker_test: install-deps non_docker_lint non_docker_compile
non_docker_test: install-deps non_docker_lint
	@echo "******* Checking if test code compiles... *************" && \
	go list ./... | grep -v /vendor/ | xargs -n1 -t -I % sh -c 'go test -c % || exit 255'
	@echo "******* Running tests... ******************************"
	go list ./... | grep -v /vendor/ | xargs -n1 -t -I % sh -c 'go test -v --cover --timeout 60s % || exit 255'

quick_cover_test:
	go tool vet -all config shared
	@DIRS="config/... shared/..." && FAILED="false" && \
	echo "gofmt -l *.go config shared" && \
	GOFMT=$$(gofmt -l *.go config shared) && \
	if [ ! -z "$$GOFMT" ]; then echo -e "\nThe following files did not pass a 'go fmt' check:\n$$GOFMT\n" && FAILED="true"; fi; \
	for codeDir in $$DIRS; do \
		echo "golint $$codeDir" && \
		LINT="$$(golint $$codeDir)" && \
		if [ ! -z "$$LINT" ]; then echo "$$LINT" && FAILED="true"; fi; \
	done && \
	if [ "$$FAILED" = "true" ]; then exit 1; else echo "ok" ;fi

	@echo "******* Checking if test code compiles... *************" && \
	go list ./... | grep -v /vendor/ | xargs -n1 -t -I % sh -c 'go test -c % || exit 255'
	@echo "******* Running tests... ******************************"
	go list ./... | grep -v /vendor/ | xargs -n1 -t -I % sh -c 'go test -v --cover --timeout 60s % || exit 255'

#REQUIRED-CI
non_docker_ci: non_docker_compile non_docker_test

clean-vendor:
	test -d vendor && rm -rf vendor || echo "vendor doesnt exist"

get-version : version
version: ## Parse version from shared/version/version.go
version:
	@echo $(VERSION)

#compile doesn't rebuild unless something changed
#REQUIRED-CI
container: compile
	set -x ;\
	docker build \
		--build-arg VERSION=${VERSION} \
		--build-arg GIT_SHA=$(GIT_SHA) \
		--tag $(IMAGE_NAME):$(GIT_SHA) . ; \
	docker tag $(IMAGE_NAME):$(GIT_SHA) $(IMAGE_NAME):$(TAG)

dev-container:    ## makes container flotilla:1.7.3-dev and installs go deps
dev-container:
	@if [ ! -e /.dockerenv ]; then \
		echo ; \
		echo ; \
		echo "------------------------------------------------" ; \
		echo "$@: Building dev container image..." ; \
		echo "------------------------------------------------" ; \
		echo ; \
		docker images | grep '$(IMAGE_NAME)' | awk '{print $$2}' | grep -q -E '^dev$$' ; \
		if [ $$? -ne 0 ]; then  \
			docker build -f Dockerfile-dev -t $(IMAGE_NAME):dev . ; \
		fi ; \
	else \
		echo ; \
		echo "------------------------------------------------" ; \
		echo "$@: Running in Docker so skipping..." ; \
		echo "------------------------------------------------" ; \
		echo ; \
		env ; \
		echo ; \
	fi

dev-clean:  ## Remove the flight-director container image
dev-clean:
	@if [ ! -e /.dockerenv ]; then \
		if $$(docker ps | grep -q "$(IMAGE_NAME):dev"); then \
			echo "You have a running dev container.  Stop it first before using dev-clean" ;\
			exit 10; \
		fi ; \
		docker images | grep '$(IMAGE_NAME)' | awk '{print $$2}' | grep -q -E '^dev$$' ; \
		if [ $$? -eq 0 ]; then  \
			docker rmi $(IMAGE_NAME):dev  ; \
		else \
			echo "No dev image" ;\
		fi ; \
	else \
		echo ; \
		echo "------------------------------------------------" ; \
		echo "$@: Running in Docker so skipping..." ; \
		echo "------------------------------------------------" ; \
		echo ; \
		env ; \
		echo ; \
	fi

#REQUIRED-CI
coverage-xml:
coverage:
	.ci/test-cover xml

#REQUIRED-CI
coveralls:
	$(MAKE) non_docker_ginko_cover
# .ci/test-cover coveralls

#REQUIRED-CI
ginkgo-cover:
	.ci/test-cover ginkgo

test-auto: ginkgo-cover
	ginkgo watch -r -cover .

# SOURCE: https://www.gnu.org/software/make/manual/html_node/Multiple-Targets.html
#REQUIRED-CI
compile lint test ci : dev-container
	build/make/run_target_in_container.sh non_docker_$@

godepgraph:
	godepgraph .

# *****************************************************
# from capcom
# install-deps:
# 	go get -u github.com/Masterminds/glide
# 	go get -u github.com/jteeuwen/go-bindata/...
# 	rm -rf vendor
# 	glide install

# install-test-deps:
# 	go get -u github.com/golang/lint/golint
# 	go get golang.org/x/tools/cmd/cover
# 	go get github.com/onsi/ginkgo/ginkgo
# 	go get github.com/onsi/gomega

# test:
# 	make install-deps install-test-deps compile
# 	go vet $$(go list ./... | grep -v /vendor/)
# 	golint $$(go list ./... | grep -v /vendor/)
# 	go test -v --timeout 5s $$(go list ./... | grep -v /vendor/)

# compile:
# 	go-bindata -pkg resources -o resources/bindata.go resources/...
# 	go build -o ./capcom
# *****************************************************

# *****************************************************
# SOURCE: https://github.com/wellington/sass/blob/master/Makefile
#
# test:
# 	go list -f '{{if len .TestGoFiles}}"go test -short {{.ImportPath}}"{{end}}' ./... | grep -v /vendor/ | xargs -L 1 sh -c

# cover:
# 	go list -f '{{if len .TestGoFiles}}"go test -cover -short {{.ImportPath}}"{{end}}' ./... | grep -v /vendor/ | xargs -L 1 sh -c
# race:
# 	go list -f '{{if len .TestGoFiles}}"go test -race -short {{.ImportPath}}"{{end}}' ./... | grep -v /vendor/ | xargs -L 1 sh -c
# *****************************************************

# SOURCE: https://github.com/lispmeister/rpi-python3/blob/534ee5ab592f0ab0cdd04a202ca492846ab12601/Makefile
exited := $(shell docker ps -a -q -f status=exited)
kill   := $(shell docker ps | grep $(container_name) | awk '{print $$1}')
# untagged := $(shell (docker images | grep "^<none>" | awk -F " " '{print $$3}'))
# dangling := $(shell docker images -f "dangling=true" -q)
# tag := $(shell docker images | grep "$(DOCKER_IMAGE_NAME)" | grep "$(DOCKER_IMAGE_VERSION)" |awk -F " " '{print $$3}')
# latest := $(shell docker images | grep "$(DOCKER_IMAGE_NAME)" | grep "latest" | awk -F " " '{print $$3}')

# clean: ## Clean old Docker images
# ifneq ($(strip $(latest)),)
# 	@echo "Removing latest $(latest) image"
# 	docker rmi "$(DOCKER_IMAGE_NAME):latest"
# endif
# ifneq ($(strip $(tag)),)
# 	@echo "Removing tag $(tag) image"
# 	docker rmi "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)"
# endif
# ifneq ($(strip $(exited)),)
# 	@echo "Cleaning exited containers: $(exited)"
# 	docker rm -v $(exited)
# endif
# ifneq ($(strip $(dangling)),)
# 	@echo "Cleaning dangling images: $(dangling)"
# 	docker rmi $(dangling)
# endif
# 	@echo 'Done cleaning.'


docker_clean:
ifneq ($(strip $(kill)),)
	@echo "Killing containers: $(kill)"
	docker kill $(kill)
endif
ifneq ($(strip $(exited)),)
	@echo "Cleaning exited containers: $(exited)"
	docker rm -v $(exited)
endif

include build/make/*.mk
