// Code generated from Flux.g4 by ANTLR 4.10. DO NOT EDIT.

package lang // Flux

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFluxVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFluxVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitFunctionCallAssignmentStatement(ctx *FunctionCallAssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitFunctionCallStatement(ctx *FunctionCallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitCommentStatement(ctx *CommentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitFunc(ctx *FuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFluxVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
