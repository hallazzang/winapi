package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	defFile     *os.File
	outFilePath string
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
	flag.Var(&FileValue{&defFile}, "d", "definition `file`")
	flag.StringVar(&outFilePath, "o", "", "output `file`")
	flag.StringVar(&packageName, "p", "winapi", "package `name`")

	flag.Parse()

	if defFile == nil {
		fmt.Println("error: definition file must be specified")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	ds, err := parse(defFile)
	if err != nil {
		panic(err)
	}

	for _, d := range ds.Definitions {
		if d.Function != nil {
			f := *d.Function
			var ps []string
			for _, p := range f.Parameters {
				ps = append(ps, fmt.Sprintf("%s %s", p.Type, p.Name))
			}
			fmt.Printf("function: %s %s(%s)\n", f.ReturnType, f.Name, strings.Join(ps, ", "))
		} else {
			t := *d.TypeDef
			if t.SrcType.SimpleType != nil {
				fmt.Printf("type: %s -> %s\n", t.DstType, *t.SrcType.SimpleType)
			} else {
				var ms []string
				for _, m := range t.SrcType.StructType.Members {
					ms = append(ms, fmt.Sprintf("\t%s %s", m.Type, m.Name))
				}
				fmt.Printf("type: %s -> struct {\n%s\n}\n", t.DstType, strings.Join(ms, "\n"))
			}
		}
	}
}
