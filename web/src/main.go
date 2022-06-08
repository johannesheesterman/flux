package main

import (
	"fmt"
	"io"
	"net/http"
	"syscall/js"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/johannesheesterman/flux/parser"
	"github.com/johannesheesterman/flux/parser/lang"
)

func runFlux() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		go func() {
			resp, _ := http.Get(args[0].String())
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))

			input := antlr.NewInputStream(string(body))
			lexer := lang.NewFluxLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := lang.NewFluxParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true
			tree := p.Prog()
			parser.NewFluxVistor().Visit(tree)
		}()

		return nil
	})
}

func main() {
	js.Global().Set("runFlux", runFlux())
	select {}
}
