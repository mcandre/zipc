// Package main provides a zipc executable for creating ZIP archives with a specified current working directory.
package main

import (
	"log"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/jhoonb/archivex"
	"github.com/mcandre/zipc"
)

// Usage is a docopt-formatted specification of this application's command line interface.
const Usage = `Usage:
  zipc [options] <archive> <source>...
  zipc -h --help
  zipc -v --version

  Arguments:
    <archive>     Archive output path
    <source>      A list of files and directories to recursively compress into the archive
  Options:
    -C <chdir>    Change the current working directory before compressing files
    -h --help     Show usage information
    -v --version  Show version information`

// main is the entrypoint for this application.
func main() {
	arguments, _ := docopt.Parse(Usage, nil, true, zipc.Version, false)

	directory, _ := arguments["-C"].(string)

	if directory != "" {
		err := os.Chdir(directory)

		if err != nil {
			log.Panic(err)
		}
	}

	archivePath, _ := arguments["<archive>"].(string)

	sourceRoots, _ := arguments["<source>"].([]string)

	archive := new(archivex.ZipFile)
	err := archive.Create(archivePath)

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := archive.Close(); err != nil {
			log.Panic(err)
		}
	}()

	for _, source := range sourceRoots {
		err := archive.AddAll(source, true)

		if err != nil {
			log.Panic(err)
		}
	}
}
