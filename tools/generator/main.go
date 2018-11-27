package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	outputFilePath = flag.String("output", "", "output file name (standard output if omitted)")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: generator [flags] [path...]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "flags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "error: no files to parse provided\n")
		flag.Usage()
	}

	src, err := ParseFiles(flag.Args())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", src)
}
