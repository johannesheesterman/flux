// Code generated from Flux.g4 by ANTLR 4.10. DO NOT EDIT.

package lang // Flux

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FluxParser.
type FluxVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FluxParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by FluxParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by FluxParser#functionCallAssignmentStatement.
	VisitFunctionCallAssignmentStatement(ctx *FunctionCallAssignmentStatementContext) interface{}

	// Visit a parse tree produced by FluxParser#functionCallStatement.
	VisitFunctionCallStatement(ctx *FunctionCallStatementContext) interface{}

	// Visit a parse tree produced by FluxParser#commentStatement.
	VisitCommentStatement(ctx *CommentStatementContext) interface{}

	// Visit a parse tree produced by FluxParser#func.
	VisitFunc(ctx *FuncContext) interface{}

	// Visit a parse tree produced by FluxParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by FluxParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by FluxParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by FluxParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
