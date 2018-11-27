package main

import (
	"bufio"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
)

type Source struct {
	PackageName string
	Functions   []*Function
	Imports     map[string]bool
}

func NewSource() *Source {
	return &Source{Imports: make(map[string]bool)}
}

func ParseFiles(fs []string) (*Source, error) {
	src := NewSource()
	src.Imports["unsafe"] = true
	for _, f := range fs {
		s, err := ParseFile(f)
		if err != nil {
			return nil, err
		}
		if src.PackageName == "" {
			src.PackageName = s.PackageName
		} else if s.PackageName != src.PackageName {
			return nil, fmt.Errorf("package name mismatch")
		}
		src.Functions = append(src.Functions, s.Functions...)
		for imp, _ := range s.Imports {
			src.Imports[imp] = true
		}
	}
	return src, nil
}

func ParseFile(path string) (*Source, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src := NewSource()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.HasPrefix(line, "//winapi") {
			continue
		}
		line = line[8:]
		if !(line[0] == ' ' || line[0] == '\t') {
			continue
		}
		line = line[1:]

		f, err := NewFunction(line)
		if err != nil {
			return nil, err
		}

		src.Functions = append(src.Functions, f)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	pkg, err := parser.ParseFile(fset, "", f, parser.PackageClauseOnly)
	if err != nil {
		return nil, err
	}
	src.PackageName = pkg.Name.Name

	return src, nil
}
