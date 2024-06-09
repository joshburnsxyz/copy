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
		return
	}
}
