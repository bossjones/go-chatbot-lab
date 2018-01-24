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
