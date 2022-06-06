package parser

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/johannesheesterman/lambda/parser/lang"
)

var standardFunctions = map[string]func([]interface{}) interface{}{
	"Process": Process,
	"Print":   Print,
}

type LambdaVisitor struct {
	lang.BaseLambdaVisitor
	values map[string]interface{}
}

func NewLambdaVistor() *LambdaVisitor {
	visitor := new(LambdaVisitor)
	visitor.values = make(map[string]interface{})
	return visitor
}

func (this *LambdaVisitor) Visit(tree antlr.Tree) {
	switch t := tree.(type) {
	case *lang.ProgContext:
		this.VisitProg(t)
	case *lang.AssignmentStatementContext:
		this.VisitAssignmentStatement(t)
	case *lang.FunctionCallAssignmentStatementContext:
		this.VisitFunctionCallAssignmentStatement(t)
	case *lang.FunctionCallStatementContext:
		this.VisitFunctionCallStatement(t)
	}
}

func (this *LambdaVisitor) VisitProg(ctx *lang.ProgContext) {
	for _, c := range ctx.GetChildren() {
		this.Visit(c)
	}
}

func (this *LambdaVisitor) VisitAssignmentStatement(ctx *lang.AssignmentStatementContext) {
	name := ctx.ID().GetText()
	value := ctx.Value()
	this.values[name] = this.VisitValue(value.(*lang.ValueContext))
}

func (this *LambdaVisitor) VisitFunctionCallAssignmentStatement(ctx *lang.FunctionCallAssignmentStatementContext) {
	name := ctx.ID().GetText()
	this.values[name] = this.VisitFunc(ctx.Func().(*lang.FuncContext))
}

func (this *LambdaVisitor) VisitFunctionCallStatement(ctx *lang.FunctionCallStatementContext) {
	this.VisitFunc(ctx.Func().(*lang.FuncContext))
}

func (this *LambdaVisitor) VisitFunc(ctx *lang.FuncContext) interface{} {
	args := make([]interface{}, len(ctx.AllValue()))
	for i, c := range ctx.AllValue() {
		args[i] = this.VisitValue(c.(*lang.ValueContext))
	}

	functionName := ctx.KEY().GetText()
	function := standardFunctions[functionName]
	return function(args)
}

func (this *LambdaVisitor) VisitValue(ctx *lang.ValueContext) interface{} {
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
		return this.VisitArray(ctx.Array().(*lang.ArrayContext))
	}
	if (ctx.Obj()) != nil {
		return this.VisitObj(ctx.Obj().(*lang.ObjContext))
	}
	if (ctx.ID()) != nil {
		return this.values[ctx.ID().GetText()]
	}

	return nil
}

func (this *LambdaVisitor) VisitArray(ctx *lang.ArrayContext) []interface{} {
	var values []interface{}
	for _, c := range ctx.AllValue() {
		values = append(values, this.VisitValue(c.(*lang.ValueContext)))
	}
	return values
}

func (this *LambdaVisitor) VisitObj(ctx *lang.ObjContext) map[string]interface{} {
	values := make(map[string]interface{})
	for _, c := range ctx.AllPair() {
		pair := c.(*lang.PairContext)
		values[pair.KEY().GetText()] = this.VisitValue(pair.Value().(*lang.ValueContext))
	}
	return values
}
