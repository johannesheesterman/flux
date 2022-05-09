package main

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/johannesheesterman/lambda/parser"
)

type LambdaVisitor struct {
	parser.BaseLambdaVisitor
	values map[string]interface{}
}

func NewLambdaVistor() *LambdaVisitor {
	visitor := new(LambdaVisitor)
	visitor.values = make(map[string]interface{})
	return visitor
}

func (this *LambdaVisitor) Visit(tree antlr.Tree) {
	switch t := tree.(type) {
	case *parser.ProgContext:
		this.VisitProg(t)
	case *parser.AssignmentStatementContext:
		this.VisitAssignmentStatement(t)
	}
}

func (this *LambdaVisitor) VisitProg(ctx *parser.ProgContext) {
	for _, c := range ctx.GetChildren() {
		this.Visit(c)
	}
}

func (this *LambdaVisitor) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) {
	name := ctx.Id().GetText()
	value := ctx.Value()
	this.values[name] = this.VisitValue(value.(*parser.ValueContext))
	fmt.Println("VisitAssignmentStatement", name, this.values[name])
}

func (this *LambdaVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	if (ctx.STRING()) != nil {
		return ctx.STRING().GetText()
	}
	if ctx.NUMBER() != nil {
		f, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
		return f
	}
	if ctx.BOOLEAN() != nil {
		return ctx.BOOLEAN().GetText() == "true"
	}
	if (ctx.Array()) != nil {
		return this.VisitArray(ctx.Array().(*parser.ArrayContext))
	}

	return nil
}

func (this *LambdaVisitor) VisitArray(ctx *parser.ArrayContext) []interface{} {
	var values []interface{}
	for _, c := range ctx.AllValue() {
		values = append(values, this.VisitValue(c.(*parser.ValueContext)))
	}
	return values
}
