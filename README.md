# zipc - like zip, but with a -chdir option

# EXAMPLE

```console
$ mkdir -p bin

$ zipc -root bin bin/zipc-0.0.3.zip zipc-0.0.3

$ unzip -l bin/zipc-0.0.3.zip
Archive:  bin/zipc-0.0.3.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  2021-08-17 15:21   zipc-0.0.3/darwin/
        0  2021-08-17 15:21   zipc-0.0.3/darwin/386/
  2375508  2021-08-17 15:21   zipc-0.0.3/darwin/386/zipc
...

$ zipc -help
  -root string
        Resolve source paths in terms of a top level root directory
  -help
        Show usage information
  -version
        Show version information
```

# DOWNLOAD

https://github.com/mcandre/zipc/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/zipc

# RUNTIME REQUIREMENTS

(None)

## Recommended

* [unzip](https://linux.die.net/man/1/unzip)

# CONTRIBUTING

For more information on developing accio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# CREDITS

Shout out to [jhoonb/archivex](https://github.com/jhoonb/archivex) for simplifying recursive archiving!

# SEE ALSO

* [go-chop](https://github.com/mcandre/go-chop)
* [go-hextime](https://github.com/mcandre/go-hextime)
* [go-ios7crypt](https://github.com/mcandre/go-ios7crypt)
* [go-rotate](https://github.com/mcandre/go-rotate)
* [go-swatch](https://github.com/mcandre/go-swatch)
* [go/forkbomb](https://github.com/mcandre/forkbombs/tree/master/go/forkbomb)
* [karp](https://github.com/mcandre/karp)
