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
SOURCES    := $(shell find . \( -name vendor \) -prune -o  -name '*.go')

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

install-deps: install-tools get-deps

get-deps:
	glide install

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
	# @which go-bindata || go get -u github.com/jteeuwen/go-bindata/...

force-vendor:
	rm -fv vendor/github.com/bossjones/go-chatbot-lab
	mkdir -p vendor/github.com/bossjones/go-chatbot-lab
	cp -af . vendor/github.com/bossjones/go-chatbot-lab

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
test:
	go test $(glide nv)

#REQUIRED-CI
# FIXME: Skip this for now cause of non_docker_compile (12/6/2017)
# non_docker_compile: install-deps build

non_docker_lint: install-deps
	go tool vet -all config shared log
	@DIRS="config/... shared/... log/..." && FAILED="false" && \
	echo "gofmt -l *.go config shared log" && \
	GOFMT=$$(gofmt -l *.go config shared log) && \
	if [ ! -z "$$GOFMT" ]; then echo -e "\nThe following files did not pass a 'go fmt' check:\n$$GOFMT\n" && FAILED="true"; fi; \
	for codeDir in $$DIRS; do \
		echo "golint $$codeDir" && \
		LINT="$$(golint $$codeDir)" && \
		if [ ! -z "$$LINT" ]; then echo "$$LINT" && FAILED="true"; fi; \
	done && \
	if [ "$$FAILED" = "true" ]; then exit 1; else echo "ok" ;fi


#REQUIRED-CI
# FIXME: Skip this for now cause of non_docker_compile (12/6/2017)
# non_docker_test: install-deps non_docker_lint non_docker_compile
non_docker_test: install-deps non_docker_lint
	@echo "******* Checking if test code compiles... *************" && \
	go list ./... | grep -v /vendor/ | xargs -n1 -t -I % sh -c 'go test -c % || exit 255'
	@echo "******* Running tests... ******************************"
	go list ./... | grep -v /vendor/ | xargs -n1 -t -I % sh -c 'go test -v --cover --timeout 60s % || exit 255'

#REQUIRED-CI
non_docker_ci: non_docker_compile non_docker_test

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
