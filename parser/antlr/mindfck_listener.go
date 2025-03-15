// Code generated from parser/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

	// EnterArrayDeclaration is called when entering the arrayDeclaration production.
	EnterArrayDeclaration(c *ArrayDeclarationContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterIfConditional is called when entering the ifConditional production.
	EnterIfConditional(c *IfConditionalContext)

	// EnterWhileLoop is called when entering the whileLoop production.
	EnterWhileLoop(c *WhileLoopContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterDebug is called when entering the debug production.
	EnterDebug(c *DebugContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterArrayAccess is called when entering the arrayAccess production.
	EnterArrayAccess(c *ArrayAccessContext)

	// EnterArrayItem is called when entering the arrayItem production.
	EnterArrayItem(c *ArrayItemContext)

	// EnterArrayIndex is called when entering the arrayIndex production.
	EnterArrayIndex(c *ArrayIndexContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitArrayDeclaration is called when exiting the arrayDeclaration production.
	ExitArrayDeclaration(c *ArrayDeclarationContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitIfConditional is called when exiting the ifConditional production.
	ExitIfConditional(c *IfConditionalContext)

	// ExitWhileLoop is called when exiting the whileLoop production.
	ExitWhileLoop(c *WhileLoopContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitDebug is called when exiting the debug production.
	ExitDebug(c *DebugContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitArrayAccess is called when exiting the arrayAccess production.
	ExitArrayAccess(c *ArrayAccessContext)

	// ExitArrayItem is called when exiting the arrayItem production.
	ExitArrayItem(c *ArrayItemContext)

	// ExitArrayIndex is called when exiting the arrayIndex production.
	ExitArrayIndex(c *ArrayIndexContext)
}
