#!/usr/bin/env bash

#   either ssh key or agent is needed to pull adobe-platform sources from git
#   this supplies to methods

set -o errexit
set -o pipefail

TARGET="$1"
SSH1=""
SSH2=""
SHA=${SHA:-"$(git rev-parse HEAD)"}

if [ ! -e /.dockerenv ]; then
    echo
    echo
    echo "-----------------------------------------------------"
    echo "Running target \"$TARGET\" inside Docker container..."
    echo "-----------------------------------------------------"
    echo
    set -x
    docker run -i --rm $SSH1 $SSH2 $AWS_ENV_VAR_OPTS \
        --name=go_chatbot_lab_make_docker_$TARGET \
        -e sha=$SHA \
        -v $PWD:/go/src/github.com/bossjones/go-chatbot-lab \
        -w /go/src/github.com/bossjones/go-chatbot-lab \
        bossjones/go-chatbot-lab:dev \
        make $TARGET
else
    make $TARGET
fi

