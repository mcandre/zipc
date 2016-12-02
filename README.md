# zipc - like zip, but with a -C <chdir> option

# EXAMPLE

```
$ zipc -C bin zipc-0.0.1.zip zipc-0.0.1

$ unzip -l bin/zipc-0.0.1.zip
Archive:  bin/zipc-0.0.1.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  12-01-2016 22:25   zipc-0.0.1/darwin/
        0  12-01-2016 22:25   zipc-0.0.1/darwin/386/
  2020040  12-01-2016 22:25   zipc-0.0.1/darwin/386/zipc
...

$ zipc -h
Usage:
  zipc [-C <chdir>] <archive> <source>...
  zipc -h --help
  zipc -v --version

  Arguments:
    <archive>     Archive output path
    <source>      A list of files and directories to recursively compress into the archive
  Options:
    -C <chdir>    Change the current working directory before compressing files
    -h --help     Show usage information
    -v --version  Show version information
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

https://github.com/mcandre/goport/releases

# REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/zipc/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/zipc.git $GOPATH/src/github.com/mcandre/zipc
$ cd $GOPATH/src/github.com/mcandre/zipc
$ git submodule update --init --recursive
$ sh -c 'cd cmd/zipc && go install'
```

# TEST

```
$ make integration-test
```

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.

# CREDITS

Shout out to [jhoonb/archivex](https://github.com/jhoonb/archivex) for simplifying recursive archiving!
