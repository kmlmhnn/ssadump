package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: ssadump filename function\n")
	}
	filename := os.Args[1]
	function := os.Args[2]

	fileinfo, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	fset.AddFile(filename, 1, int(fileinfo.Size()))

	astRoot, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	pkg, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()},
		fset,
		types.NewPackage(filename, ""),
		[]*ast.File{astRoot},
		ssa.SanityCheckFunctions,
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range pkg.Func(function).Blocks {
		fmt.Printf("block %d:\n", b.Index)
		for _, instr := range b.Instrs {
			val, ok := instr.(ssa.Value)
			switch ok {
			case true:
				fmt.Printf("%-20T %5s  :=  %s\n", val, val.Name(), val.String())
			case false:
				fmt.Printf("%-20T %5s      %s\n", instr, "N/A", instr.String())
			}

		}
		fmt.Println()
	}

}
