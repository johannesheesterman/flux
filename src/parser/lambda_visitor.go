// Code generated from Lambda.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Lambda

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LambdaParser.
type LambdaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LambdaParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by LambdaParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by LambdaParser#variableDeclarationStatement.
	VisitVariableDeclarationStatement(ctx *VariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by LambdaParser#commentStatement.
	VisitCommentStatement(ctx *CommentStatementContext) interface{}

	// Visit a parse tree produced by LambdaParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by LambdaParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by LambdaParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by LambdaParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
