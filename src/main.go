package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/johannesheesterman/lambda/parser"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "test.lda")
	}
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewLambdaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLambdaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	NewLambdaVistor().Visit(tree)
}
