package main

import (
	"fmt"
	"strings"
)

type FunctionReference struct {
	DllName      string
	FunctionName string
}

type Function struct {
	Reference  *FunctionReference
	Name       string
	Parameters []*Variable
	Returns    []*Variable
}

func NewFunction(s string) (*Function, error) {
	f := new(Function)
	return f, nil
}

func (f *Function) BodyString() string {
	return ""
}

func (f *Function) ParametersString() string {
	var ps []string
	for _, p := range f.Parameters {
		ps = append(ps, p.String())
	}
	return strings.Join(ps, ", ")
}

func (f *Function) ReturnsString() string {
	switch len(f.Returns) {
	case 0:
		return ""
	case 1:
		return f.Returns[0].String()
	default:
		var rs []string
		for _, r := range f.Returns {
			rs = append(rs, r.String())
		}
		return "(" + strings.Join(rs, ", ") + ")"
	}
}

func (f *Function) String() string {
	s := fmt.Sprintf("func %s(%s)", f.Name, f.ParametersString())
	rs := f.ReturnsString()
	if len(rs) > 0 {
		s += " " + rs
	}
	return s
}

type Variable struct {
	Name string
	Type string
}

func (v *Variable) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Type)
}
