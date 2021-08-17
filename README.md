# zipc - like zip, but with a -chdir option

# EXAMPLE

```console
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

# CONTRIBUTING

For more information on developing accio itself, see [DEVELOPMENT.md](DEVELOPMENT.md).

# CREDITS

Shout out to [jhoonb/archivex](https://github.com/jhoonb/archivex) for simplifying recursive archiving!
