// Code generated from parser/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindfck

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by mindfckParser.
type mindfckVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by mindfckParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by mindfckParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by mindfckParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by mindfckParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by mindfckParser#print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by mindfckParser#ifConditional.
	VisitIfConditional(ctx *IfConditionalContext) interface{}

	// Visit a parse tree produced by mindfckParser#whileLoop.
	VisitWhileLoop(ctx *WhileLoopContext) interface{}

	// Visit a parse tree produced by mindfckParser#read.
	VisitRead(ctx *ReadContext) interface{}

	// Visit a parse tree produced by mindfckParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by mindfckParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by mindfckParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by mindfckParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
