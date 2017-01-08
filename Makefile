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
	  go build -ldflags="-s -w" -o bin/tar-append-linux64
	@docker run --rm \
	  -v $(PWD):/go/src/github.com/allingeek/tar-append \
	  -v $(PWD)/bin:/go/bin \
	  -w /go/src/github.com/allingeek/tar-append \
	  -e GOOS=darwin \
	  -e GOARCH=amd64 \
	  golang:1.7 \
	  go build -ldflags="-s -w" -o bin/tar-append-darwin64
upx: build
	@docker run --rm \
	  -v $(PWD)/bin:/input \
	  -w /input \
	  allingeek/upx:latest \
	  --brute -k tar-append-linux64
	@mv ./bin/tar-append-linux64 ./bin/tar-append-linux64-upx
	@mv ./bin/tar-append-linux64.~ ./bin/tar-append-linux64
	@docker run --rm \
	  -v $(PWD)/bin:/input \
	  -w /input \
	  allingeek/upx:latest \
	  --brute -k tar-append-darwin64
	@mv ./bin/tar-append-darwin64 ./bin/tar-append-darwin64-upx
	@mv ./bin/tar-append-darwin64.~ ./bin/tar-append-darwin64
