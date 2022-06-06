package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/johannesheesterman/lambda/parser"
	"github.com/johannesheesterman/lambda/parser/lang"
)

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "test.lda")
	}
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := lang.NewLambdaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := lang.NewLambdaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	parser.NewLambdaVistor().Visit(tree)
}
