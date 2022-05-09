package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/johannesheesterman/lambda/parser"
)

type LambdaVisitor struct {
	parser.BaseLambdaVisitor
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
	fmt.Println("VisitAssignmentStatement", ctx.ID().GetText())
}
