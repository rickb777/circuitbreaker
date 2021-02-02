#!/bin/bash -e
cd "$(dirname $0)"
PATH=$HOME/go/bin:$PATH
unset GOPATH
export GO111MODULE=on
export GOARCH=${1}

function v
{
  echo
  echo $@
  $@
}

if ! type -p goveralls; then
  v go install github.com/mattn/goveralls
fi

if ! type -p shadow; then
  v go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
fi

if ! type -p goreturns; then
  v go install github.com/sqs/goreturns
fi

echo circuitbreaker...
v go test -v -covermode=count -coverprofile=circuitbreaker.out .
v go tool cover -func=circuitbreaker.out
#[ -z "$COVERALLS_TOKEN" ] || goveralls -coverprofile=circuitbreaker.out -service=travis-ci -repotoken $COVERALLS_TOKEN

v goreturns -l -w *.go

v go vet ./...

v shadow ./...
