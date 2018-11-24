package main

import (
	"io"

	"github.com/alecthomas/participle"
)

type Definitions struct {
	Definitions []*Definition `( @@ )*`
}

type Definition struct {
	TypeDef  *TypeDef  `  "typedef" @@`
	Function *Function `| @@`
}

type Function struct {
	ReturnType string       `@Ident`
	Name       string       `@Ident "("`
	Parameters []*Parameter `( @@ ( "," @@ )* )*`
	_          string       `")" ";"`
}

type Parameter struct {
	Type string `@Ident`
	Name string `@Ident`
}

type TypeDef struct {
	SrcType *Type  `@@`
	DstType string `@Ident`
	_       string `";"`
}

type Type struct {
	StructType *StructType `"struct" @@`
	SimpleType *string     `| @Ident`
}

type StructType struct {
	_       string              `"{"`
	Members []*StructTypeMember `( @@ )*`
	_       string              `"}"`
}

type StructTypeMember struct {
	Type string `@Ident`
	Name string `@Ident`
	_    string `";"`
}

func parse(r io.Reader) (*Definitions, error) {
	d := new(Definitions)

	parser, err := participle.Build(d, participle.Elide("Comment"))
	if err != nil {
		return nil, err
	}

	if err := parser.Parse(r, d); err != nil {
		return nil, err
	}

	return d, err
}
