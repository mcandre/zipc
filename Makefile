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

shlint:
	find . \( -wholename '*/node_modules*' \) -prune -o -type f \( -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs shlint

checkbashisms:
	find . \( -wholename '*/node_modules*' \) -prune -o -type f \( -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs checkbashisms -n -p

shellcheck:
	find . \( -wholename '*/node_modules*' \) -prune -o -type f \( -name '*.sh' -o -name '*.bashrc*' -o -name '.*profile*' -o -name '*.envrc*' \) -print | xargs shellcheck

lint: govet gofmt goimport shlint checkbashisms shellcheck

clean: clean-ports

clean-ports:
	rm -rf bin
