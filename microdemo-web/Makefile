.PHONY: build test clean

_APP="microdemo-web"

build: clean
	@CGO_ENABLED=0 go build -o $(_APP) *.go
	@./$(_APP) --version

test:
	@go test -v ./... -cover

clean:
	@rm -f $(_APP)
