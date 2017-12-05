.PHONY: build build-alpine clean test help default

# SOURCES:=$(shell find . \( -name vendor \) -prune -o  -name '*.go')
# RESOURCES:=$(shell find ./resources -type f  | grep -v resources/bindata.go )
# FD_VERSION  = $(shell awk -F "\"" '/var Version/ { print $$2 }' shared/version/version.go)

BIN_NAME=go-chatbot-lab

VERSION    := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_SHA  = $(shell git rev-parse HEAD)
GIT_DIRTY   = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_BRANCH  = $(shell git rev-parse --abbrev-ref HEAD)
IMAGE_NAME := "bossjones/go-chatbot-lab"

default: test

help:
	@echo 'Management commands for go-chatbot-lab:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs glide install, mostly used for ci.'
	@echo '    make build-alpine    Compile optimized for alpine linux.'
	@echo '    make package         Build final docker image with just the go binary inside'
	@echo '    make tag             Tag image created by package with latest, git commit and version'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make push            Push tagged images to registry'
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X main.GitCommit=${GIT_SHA}${GIT_DIRTY} -X main.VersionPrerelease=DEV" -o bin/${BIN_NAME}

get-deps:
	glide install

install-tools:
# @which golint || go get -u github.com/golang/lint/golint
# @which cover || go get golang.org/x/tools/cmd/cover
# @test -d $$GOPATH/github.com/go-ini/ini || go get github.com/go-ini/ini
# @test -d $$GOPATH/github.com/jmespath/go-jmespath ||  go get github.com/jmespath/go-jmespath
# @which ginkgo || go get github.com/onsi/ginkgo/ginkgo
# @which gomega || go get github.com/onsi/gomega
# @which gomock || go get github.com/golang/mock/gomock
# @which mockgen || go get github.com/golang/mock/mockgen
	@which glide || go get github.com/Masterminds/glide
# @which go-bindata || go get -u github.com/jteeuwen/go-bindata/...

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

test:
	go test $(glide nv)

