// Package main provides a zipc executable for creating ZIP archives with a specified current working directory.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mcandre/zipc"
)

var flagRoot = flag.String("root", "", "Resolve source paths in terms of a top level root directory")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

// main is the entrypoint for this application.
func main() {
	flag.Parse()

	switch {
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(1)
	case *flagVersion:
		fmt.Println(zipc.Version)
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) < 2 {
		log.Panic("An archive path and at least one source path must be supplied, following any named flags")
	}

	if err := zipc.Compress(args[0], args[1:], *flagRoot); err != nil {
		log.Panic(err)
	}
}
