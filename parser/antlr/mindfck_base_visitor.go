// Code generated from parser/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindfck

import "github.com/antlr4-go/antlr/v4"

type BasemindfckVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasemindfckVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitPrint(ctx *PrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitIfConditional(ctx *IfConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitWhileLoop(ctx *WhileLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitRead(ctx *ReadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasemindfckVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
