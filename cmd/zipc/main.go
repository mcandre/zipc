// Package main provides a zipc executable for creating ZIP archives with a specified current working directory.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/jhoonb/archivex"
	"github.com/mcandre/zipc"
)

var flagDirectory = flag.String("chdir", "", "Change the current working directory before compressing files")
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

	archivePath, sourceRoots := args[0], args[1:]

	root := *flagDirectory

	if root != "" {
		for i, sourceRoot := range sourceRoots {
			sourceRoots[i] = path.Join(root, sourceRoot)
		}
	}

	archive := new(archivex.ZipFile)

	if err := archive.Create(archivePath); err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := archive.Close(); err != nil {
			log.Panic(err)
		}
	}()

	for _, source := range sourceRoots {
		if err := archive.AddAll(source, true); err != nil {
			log.Panic(err)
		}
	}
}
