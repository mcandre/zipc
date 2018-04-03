# zipc - like zip, but with a -chdir option

# EXAMPLE

```
$ mkdir -p bin

$ zipc -chdir bin zipc-0.0.2.zip zipc-0.0.2

$ unzip -l bin/zipc-0.0.2.zip
Archive:  bin/zipc-0.0.2.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  12-01-2016 22:25   zipc-0.0.2/darwin/
        0  12-01-2016 22:25   zipc-0.0.2/darwin/386/
  2020040  12-01-2016 22:25   zipc-0.0.2/darwin/386/zipc
...

$ zipc -help
  -chdir string
        Change the current working directory before compressing files
  -help
        Show usage information
  -version
        Show version information
```

More examples:

* [go-chop](https://github.com/mcandre/go-chop)
* [go-hextime](https://github.com/mcandre/go-hextime)
* [go-ios7crypt](https://github.com/mcandre/go-ios7crypt)
* [go-rotate](https://github.com/mcandre/go-rotate)
* [go-swatch](https://github.com/mcandre/go-swatch)
* [go/forkbomb](https://github.com/mcandre/forkbombs/tree/master/go/forkbomb)
* [karp](https://github.com/mcandre/karp)

# DOWNLOAD

https://github.com/mcandre/zipc/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/zipc

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.9+
* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/zipc/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p "$GOPATH/src/github.com/mcandre"
$ git clone https://github.com/mcandre/zipc.git "$GOPATH/src/github.com/mcandre/zipc"
$ cd "$GOPATH/src/github.com/mcandre/zipc"
$ git submodule update --init --recursive
$ go install ./...
```

# TEST

```
$ mage test
```

# PORT

```
$ mage port
```

# LINT

Keep the code tidy:

```
$ mage lint
```

# CREDITS

Shout out to [jhoonb/archivex](https://github.com/jhoonb/archivex) for simplifying recursive archiving!
