package main

import (
	"io"

	"github.com/alecthomas/participle"
)

type Functions struct {
	Functions []*Function `{ @@ }`
}

type Function struct {
	ReturnType string       `@Ident`
	Name       string       `@Ident "("`
	Parameters []*Parameter `{ @@ { "," @@ } }`
	_          string       `")" ";"`
}

type Parameter struct {
	Type string `@Ident`
	Name string `@Ident`
}

func parseFunctions(r io.Reader) (*Functions, error) {
	fs := new(Functions)

	parser, err := participle.Build(fs, participle.Elide("Comment"))
	if err != nil {
		return nil, err
	}

	if err := parser.Parse(r, fs); err != nil {
		return nil, err
	}

	return fs, nil
}
