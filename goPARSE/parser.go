package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage:\n\t%s [files]\n", os.Args[0])
		os.Exit(1)
	}

	var v visitor // This object what helpes us traverse the AST (hence "visitor")
	fs := token.NewFileSet()

	for _, arg := range os.Args[1:] { //iterating over the files
		f, err := parser.ParseFile(fs, arg, nil, parser.AllErrors)
		if err != nil {
			log.Printf("Could not parse %s: %v", arg, err)
			continue
		}

		ast.Walk(v, f)
	}
	fmt.Println(v.functions)
}

type visitor struct {
	functions []string
}

// @Override the Visit call
func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {
	case *ast.FuncDecl:
		v.functions = append(v.functions, d.Name.Name)
		fmt.Printf("%s\n", d.Name.Name)
	}
	return v
}
