// Code generated from Lambda.g4 by ANTLR 4.10. DO NOT EDIT.

package lang // Lambda

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLambdaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLambdaVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitFunctionCallAssignmentStatement(ctx *FunctionCallAssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitFunctionCallStatement(ctx *FunctionCallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitCommentStatement(ctx *CommentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLambdaVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
