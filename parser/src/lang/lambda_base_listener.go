// Code generated from Flux.g4 by ANTLR 4.10. DO NOT EDIT.

package lang // Flux

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFluxListener is a complete listener for a parse tree produced by FluxParser.
type BaseFluxListener struct{}

var _ FluxListener = &BaseFluxListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFluxListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFluxListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFluxListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFluxListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseFluxListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseFluxListener) ExitProg(ctx *ProgContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseFluxListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseFluxListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterFunctionCallAssignmentStatement is called when production functionCallAssignmentStatement is entered.
func (s *BaseFluxListener) EnterFunctionCallAssignmentStatement(ctx *FunctionCallAssignmentStatementContext) {
}

// ExitFunctionCallAssignmentStatement is called when production functionCallAssignmentStatement is exited.
func (s *BaseFluxListener) ExitFunctionCallAssignmentStatement(ctx *FunctionCallAssignmentStatementContext) {
}

// EnterFunctionCallStatement is called when production functionCallStatement is entered.
func (s *BaseFluxListener) EnterFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// ExitFunctionCallStatement is called when production functionCallStatement is exited.
func (s *BaseFluxListener) ExitFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// EnterCommentStatement is called when production commentStatement is entered.
func (s *BaseFluxListener) EnterCommentStatement(ctx *CommentStatementContext) {}

// ExitCommentStatement is called when production commentStatement is exited.
func (s *BaseFluxListener) ExitCommentStatement(ctx *CommentStatementContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseFluxListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseFluxListener) ExitFunc(ctx *FuncContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseFluxListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseFluxListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseFluxListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseFluxListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseFluxListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseFluxListener) ExitArray(ctx *ArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFluxListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFluxListener) ExitValue(ctx *ValueContext) {}
