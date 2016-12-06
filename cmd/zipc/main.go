package main

import (
	"os"

	"github.com/docopt/docopt-go"
	"github.com/jhoonb/archivex"
	"github.com/mcandre/zipc"
)

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

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, zipc.Version, false)

	if err != nil {
		panic(Usage)
	}

	directory, _ := arguments["-C"].(string)

	if directory != "" {
		os.Chdir(directory)
	}

	archivePath, _ := arguments["<archive>"].(string)

	sourceRoots, _ := arguments["<source>"].([]string)

	archive := new(archivex.ZipFile)
	archive.Create(archivePath)
	defer archive.Close()

	for _, source := range sourceRoots {
		archive.AddAll(source, true)
	}
}
