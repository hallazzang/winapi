package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	sigFile     *os.File
	typeFile    *os.File
	packageName string
)

// FileValue implements flag.Value interface
type FileValue struct {
	File **os.File
}

func (fv *FileValue) String() string {
	if fv.File != nil && *fv.File != nil {
		return (*fv.File).Name()
	}
	return ""
}

func (fv *FileValue) Set(s string) error {
	if f, err := os.Open(s); err != nil {
		return err
	} else {
		*fv.File = f
	}
	return nil
}

func init() {
	flag.Var(&FileValue{&sigFile}, "s", "function signature `file`")
	flag.Var(&FileValue{&typeFile}, "t", "type definition `file`")
	flag.StringVar(&packageName, "p", "winapi", "package `name`")

	flag.Parse()

	if sigFile == nil && typeFile == nil {
		flag.Usage()
		os.Exit(1)
	} else if sigFile == nil {
		fmt.Println("function signature file must be provided with -s option")
		os.Exit(1)
	} else if typeFile == nil {
		fmt.Println("type definition file must be provided with -t option")
		os.Exit(1)
	}
}

func main() {
	fmt.Println(sigFile, typeFile, packageName)
}
