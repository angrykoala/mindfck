// Code generated from parser/antlr/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindfck

import "github.com/antlr4-go/antlr/v4"

// mindfckListener is a complete listener for a parse tree produced by mindfckParser.
type mindfckListener interface {
	antlr.ParseTreeListener

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterIfConditional is called when entering the ifConditional production.
	EnterIfConditional(c *IfConditionalContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitIfConditional is called when exiting the ifConditional production.
	ExitIfConditional(c *IfConditionalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}
