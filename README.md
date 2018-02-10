# go-chatbot-lab

Development only. I'm trying to learnsome golang! First project, building a chatbot!

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/go-chatbot-lab
```

### Testing

``make test``

### Thanks

Cookiecutter!

[cookiecutter-golang](https://github.com/lacion/cookiecutter-golang/tree/master)

### Example Projects Using Same Cookiecutter

[iothub](https://github.com/lacion/iothub)

### Problems

- [Editor not looking at vendor folder it is ln -s from other path VSCode](https://github.com/Microsoft/vscode-go/issues/1327)


### Basic Golang gotyas

You go code needs to exist in `$GOPATH/src`

Eg.

```

 |2.2.3|    dev7-behance-1484 in ~/dev/go_workspace/src/github.com/bossjones
○ → code go-chatbot-lab/
```

# mockgen

```
#!/bin/bash -e

# source: https://github.com/rafrombrc/gomock/blob/master/update_mocks.sh

mockgen github.com/rafrombrc/gomock/gomock Matcher \
  > gomock/mock_matcher/mock_matcher.go
mockgen github.com/rafrombrc/gomock/sample Index,Embed,Embedded \
  > sample/mock_user/mock_user.go
gofmt -w gomock/mock_matcher/mock_matcher.go sample/mock_user/mock_user.go

echo >&2 "OK"
```

# A getting started guide for Go newcomers

https://github.com/alco/gostart


## Server code borrowed from coolspeed/century project

It's just a simple chatbot i'm building to help teach me how golang works. Actual server code borrowed from coolspeed/century project! Will add more on top of that.
## Feature

* High throughput
* High concurrency
* (Automatic) High scalability, especially on many-core computers. (Think of 64-core computers, as much as 4-core ones.)

## Detailed Information

You can find an even simpler chat server on:

[https://gist.github.com/drewolson/3950226](https://gist.github.com/drewolson/3950226)

(In fact I started my scratch from that.)

## Build and Run

1) First, you need to install `golang`, of course:

Download it from [https://golang.org/dl/](https://golang.org/dl/), or install go via your package management tool:

For Ubuntu:

```
sudo apt-get install golang
```

For Mac OS X:

```
brew install go
```

2) Now, just build.

`cd` into the repo directory.

To build the server, execute:

```
make build
```

3) Run

3.1 Run the chat server:

```
./bin/go-chatbot-lab
```

3.2 Run the chat client:

`Client`: You can use `telnet` as the client:

```
telnet localhost 6666
```

type anything.
