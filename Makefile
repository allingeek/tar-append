# needs Docker
PWD := $(shell pwd)

build:
	@docker run --rm \
	  -v $(PWD):/go/src/github.com/allingeek/tar-append \
	  -v $(PWD)/bin:/go/bin \
	  -w /go/src/github.com/allingeek/tar-append \
	  -e GOOS=linux \
	  -e GOARCH=amd64 \
	  golang:1.7 \
	  go build -o bin/tar-append-linux64
	@docker run --rm \
	  -v $(PWD):/go/src/github.com/allingeek/tar-append \
	  -v $(PWD)/bin:/go/bin \
	  -w /go/src/github.com/allingeek/tar-append \
	  -e GOOS=darwin \
	  -e GOARCH=amd64 \
	  golang:1.7 \
	  go build -o bin/tar-append-darwin64
