.PHONY: build-linux build-mac

LINUX=env GOOS=linux GOARCH=amd64 go build -v
MAC=env GOOS=darwin GOARCH=amd64 go build -v

build-linux:
	$(LINUX) -o bin/ccwc-linux cmd/ccwc/*.go

build-mac:
	$(MAC) -o bin/ccwc-mac cmd/ccwc/*.go
