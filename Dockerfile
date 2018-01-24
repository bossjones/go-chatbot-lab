# Build Stage
FROM golang:1.9.2-stretch AS build-stage

LABEL app="build-go-chatbot-lab"
LABEL REPO="https://github.com/bossjones/go-chatbot-lab"

# ENV GOROOT=/usr/lib/go \
#     GOPATH=/go \
#     GOBIN=/gop/bin \
#     PROJPATH=/go/src/github.com/bossjones/go-chatbot-lab

ENV GOPATH=/go \
    GOBIN=/go/bin \
    PROJPATH=/go/src/github.com/bossjones/go-chatbot-lab

# Because of https://github.com/docker/docker/issues/14914
# ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
# ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/bossjones/go-chatbot-lab
WORKDIR /go/src/github.com/bossjones/go-chatbot-lab

RUN make build

# Final Stage
FROM behance/docker-base:2.0.1

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/bossjones/go-chatbot-lab"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/go-chatbot-lab/bin

WORKDIR /opt/go-chatbot-lab/bin

COPY --from=build-stage /go/src/github.com/bossjones/go-chatbot-lab/bin/go-chatbot-lab /opt/go-chatbot-lab/bin/
RUN chmod +x /opt/go-chatbot-lab/bin/go-chatbot-lab

CMD /opt/go-chatbot-lab/bin/go-chatbot-lab
