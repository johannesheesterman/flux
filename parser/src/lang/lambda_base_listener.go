// Code generated from Lambda.g4 by ANTLR 4.10. DO NOT EDIT.

package lang // Lambda

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLambdaListener is a complete listener for a parse tree produced by LambdaParser.
type BaseLambdaListener struct{}

var _ LambdaListener = &BaseLambdaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLambdaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLambdaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLambdaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLambdaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseLambdaListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseLambdaListener) ExitProg(ctx *ProgContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseLambdaListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseLambdaListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterFunctionCallAssignmentStatement is called when production functionCallAssignmentStatement is entered.
func (s *BaseLambdaListener) EnterFunctionCallAssignmentStatement(ctx *FunctionCallAssignmentStatementContext) {
}

// ExitFunctionCallAssignmentStatement is called when production functionCallAssignmentStatement is exited.
func (s *BaseLambdaListener) ExitFunctionCallAssignmentStatement(ctx *FunctionCallAssignmentStatementContext) {
}

// EnterFunctionCallStatement is called when production functionCallStatement is entered.
func (s *BaseLambdaListener) EnterFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// ExitFunctionCallStatement is called when production functionCallStatement is exited.
func (s *BaseLambdaListener) ExitFunctionCallStatement(ctx *FunctionCallStatementContext) {}

// EnterCommentStatement is called when production commentStatement is entered.
func (s *BaseLambdaListener) EnterCommentStatement(ctx *CommentStatementContext) {}

// ExitCommentStatement is called when production commentStatement is exited.
func (s *BaseLambdaListener) ExitCommentStatement(ctx *CommentStatementContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseLambdaListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseLambdaListener) ExitFunc(ctx *FuncContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseLambdaListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseLambdaListener) ExitObj(ctx *ObjContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseLambdaListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseLambdaListener) ExitPair(ctx *PairContext) {}

// EnterArray is called when production array is entered.
func (s *BaseLambdaListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseLambdaListener) ExitArray(ctx *ArrayContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLambdaListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLambdaListener) ExitValue(ctx *ValueContext) {}
