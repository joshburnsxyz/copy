package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

var (
	showHelp     bool
	showProgress bool
	recursive    bool
	sourcePath   string
	destPath     string
)

func init() {
	flag.BoolVarP(&showHelp, "help", "h", false, "Display help message")
	flag.BoolVarP(&showProgress, "progress", "p", false, "Display progress bar")
	flag.BoolVarP(&recursive, "recursive", "R", false, "Copy directory recursively")

	// Usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] SRC DEST\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	// Display help message if asked for it
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	// if we dont get both args, print usage and exit
	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	// get source and dest paths as strings
	sourcePath = args[0]
	destPath = args[1]

	// Check source exists
	file, err := os.Stat(sourcePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err)
		os.Exit(1)
	}

	// check if recursive flag is enabled for directory sources
	if file.IsDir() {
		if !recursive {
			fmt.Fprintf(os.Stderr, "ERROR: source is a directory, use -R to copy recursively")
			os.Exit(1)
		}
		err = copyDir(sourcePath, destPath, showProgress)
	} else {
		err = copyFile(sourcePath, destPath, showProgress)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s", err)
		os.Exit(1)
	}
}

func copyFile(src string, dest string, progress bool) error {
	// TODO: implement copy file function
}

func copyDir(src string, dest string, progress bool) error {
	// TODO: implement copy file function
}
