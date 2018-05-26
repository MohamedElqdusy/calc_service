.PHONY: build doc fmt lint run test vendor_clean vendor_get vendor_update vet

GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

default: build

build: vet
	go build -v -o ./bin/server ./src/server
	go build -v -o ./bin/client ./src/client

doc:
	godoc -http=:6060 -index

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

# https://github.com/golang/lint
#	go get github.com/golang/lint/golint
lint:
	golint ./src

run: build
	./bin/server
	./bin/client

test:
	go test ./src/...

vendor_clean:
	rm -dRf ./_vendor/src

vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor go get -d -u -v \


vendor_update: vendor_get
	rm -rf `find ./_vendor/src -type d -name .git` \
	&& rm -rf `find ./_vendor/src -type d -name .hg` \
	&& rm -rf `find ./_vendor/src -type d -name .bzr` \
	&& rm -rf `find ./_vendor/src -type d -name .svn`

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./src/...
