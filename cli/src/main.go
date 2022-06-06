package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/johannesheesterman/flux/parser"
	"github.com/johannesheesterman/flux/parser/lang"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "test.flux")
	}
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := lang.NewFluxLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := lang.NewFluxParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	parser.NewFluxVistor().Visit(tree)
}
