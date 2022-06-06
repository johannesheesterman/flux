package main

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/johannesheesterman/lambda/parser"
)

var standardFunctions = map[string]func([]interface{}) interface{}{
	"Process": Process,
	"Print":   Print,
}

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
	case *parser.FunctionCallAssignmentStatementContext:
		this.VisitFunctionCallAssignmentStatement(t)
	case *parser.FunctionCallStatementContext:
		this.VisitFunctionCallStatement(t)
	}
}

func (this *LambdaVisitor) VisitProg(ctx *parser.ProgContext) {
	for _, c := range ctx.GetChildren() {
		this.Visit(c)
	}
}

func (this *LambdaVisitor) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) {
	name := ctx.ID().GetText()
	value := ctx.Value()
	this.values[name] = this.VisitValue(value.(*parser.ValueContext))
}

func (this *LambdaVisitor) VisitFunctionCallAssignmentStatement(ctx *parser.FunctionCallAssignmentStatementContext) {
	name := ctx.ID().GetText()
	this.values[name] = this.VisitFunc(ctx.Func().(*parser.FuncContext))
}

func (this *LambdaVisitor) VisitFunctionCallStatement(ctx *parser.FunctionCallStatementContext) {
	this.VisitFunc(ctx.Func().(*parser.FuncContext))
}

func (this *LambdaVisitor) VisitFunc(ctx *parser.FuncContext) interface{} {
	args := make([]interface{}, len(ctx.AllValue()))
	for i, c := range ctx.AllValue() {
		args[i] = this.VisitValue(c.(*parser.ValueContext))
	}

	functionName := ctx.KEY().GetText()
	function := standardFunctions[functionName]
	return function(args)
}

func (this *LambdaVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	if (ctx.STRING()) != nil {
		return ctx.STRING().GetText()[1 : len(ctx.STRING().GetText())-1]
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
	if (ctx.Obj()) != nil {
		return this.VisitObj(ctx.Obj().(*parser.ObjContext))
	}
	if (ctx.ID()) != nil {
		return this.values[ctx.ID().GetText()]
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

func (this *LambdaVisitor) VisitObj(ctx *parser.ObjContext) map[string]interface{} {
	values := make(map[string]interface{})
	for _, c := range ctx.AllPair() {
		pair := c.(*parser.PairContext)
		values[pair.KEY().GetText()] = this.VisitValue(pair.Value().(*parser.ValueContext))
	}
	return values
}
