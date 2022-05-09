// Code generated from Lambda.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Lambda

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LambdaListener is a complete listener for a parse tree produced by LambdaParser.
type LambdaListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterCommentStatement is called when entering the commentStatement production.
	EnterCommentStatement(c *CommentStatementContext)

	// EnterObj is called when entering the obj production.
	EnterObj(c *ObjContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitCommentStatement is called when exiting the commentStatement production.
	ExitCommentStatement(c *CommentStatementContext)

	// ExitObj is called when exiting the obj production.
	ExitObj(c *ObjContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)
}
