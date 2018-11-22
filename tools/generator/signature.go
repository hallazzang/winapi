package main

import (
	"io"

	"github.com/alecthomas/participle"
)

type Signatures struct {
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

func parseSignatures(r io.Reader) (*Signatures, error) {
	sigs := new(Signatures)

	parser, err := participle.Build(sigs, participle.Elide("Comment"))
	if err != nil {
		return nil, err
	}

	if err := parser.Parse(r, sigs); err != nil {
		return nil, err
	}

	return sigs, nil
}
