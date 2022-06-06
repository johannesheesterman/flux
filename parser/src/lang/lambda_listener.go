// Code generated from Lambda.g4 by ANTLR 4.10. DO NOT EDIT.

package lang // Lambda

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LambdaListener is a complete listener for a parse tree produced by LambdaParser.
type LambdaListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterFunctionCallAssignmentStatement is called when entering the functionCallAssignmentStatement production.
	EnterFunctionCallAssignmentStatement(c *FunctionCallAssignmentStatementContext)

	// EnterFunctionCallStatement is called when entering the functionCallStatement production.
	EnterFunctionCallStatement(c *FunctionCallStatementContext)

	// EnterCommentStatement is called when entering the commentStatement production.
	EnterCommentStatement(c *CommentStatementContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitFunctionCallAssignmentStatement is called when exiting the functionCallAssignmentStatement production.
	ExitFunctionCallAssignmentStatement(c *FunctionCallAssignmentStatementContext)

	// ExitFunctionCallStatement is called when exiting the functionCallStatement production.
	ExitFunctionCallStatement(c *FunctionCallStatementContext)

	// ExitCommentStatement is called when exiting the commentStatement production.
	ExitCommentStatement(c *CommentStatementContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
