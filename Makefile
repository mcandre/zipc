VERSION=0.0.1

.PHONY: port clean clean-ports

all: integration-test

integration-test: port

port: bin
	zipc -C bin "zipc-$(VERSION).zip" "zipc-$(VERSION)"
	unzip -l "bin/zipc-$(VERSION).zip"

bin:
	gox -output="bin/{{.Dir}}-$(VERSION)/{{.OS}}/{{.Arch}}/{{.Dir}}" ./cmd/...

govet:
	go list ./... | grep -v vendor | xargs go vet -v

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

lint: govet gofmt goimport

clean: clean-ports

clean-ports:
	rm -rf bin
