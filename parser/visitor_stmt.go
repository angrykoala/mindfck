package parser

import (
	"mindfck/env"
	"mindfck/mfast"
	mindfck "mindfck/parser/antlr"
	"mindfck/utils"
)

func (v *AstGeneratorVisitor) VisitStatement(ctx *mindfck.StatementContext) interface{} {
	if ctx.Declaration() != nil {
		return ctx.Declaration().Accept(v)
	}
	if ctx.ArrayDeclaration() != nil {
		return ctx.ArrayDeclaration().Accept(v)
	}
	if ctx.Assignment() != nil {
		return ctx.Assignment().Accept(v)
	}
	if ctx.Print_() != nil {
		return ctx.Print_().Accept(v)
	}
	if ctx.IfConditional() != nil {
		return ctx.IfConditional().Accept(v)
	}
	if ctx.WhileLoop() != nil {
		return ctx.WhileLoop().Accept(v)
	}
	if ctx.Read() != nil {
		return ctx.Read().Accept(v)
	}
	if ctx.Debug() != nil {
		return ctx.Debug().Accept(v)
	}

	panic("Spanish Inquisition (unexpected)")
}

func (v *AstGeneratorVisitor) VisitDeclaration(ctx *mindfck.DeclarationContext) interface{} {
	var varType env.VarType
	if ctx.INT() != nil {
		varType = env.INT
	} else if ctx.BYTE() != nil {
		varType = env.BYTE
	} else {
		panic("invalid type in declaration")
	}

	var assign *mfast.Assign
	if ctx.EQUALS() != nil {
		// Declaration with assignment
		expr := ctx.Expression().Accept(v).(mfast.Expr)
		assign = &mfast.Assign{
			To:   ctx.Identifier().GetText(),
			From: expr,
		}
	}

	return &mfast.Declare{
		Label:   ctx.Identifier().IDENTIFIER().GetText(),
		VarType: varType,
		Assign:  assign,
	}
}

func (v *AstGeneratorVisitor) VisitArrayDeclaration(ctx *mindfck.ArrayDeclarationContext) interface{} {
	var assign *mfast.Assign
	if ctx.EQUALS() != nil {
		// Declaration with assignment
		expr := ctx.Expression().Accept(v).(mfast.Expr)

		assign = &mfast.Assign{
			To:   ctx.Identifier().GetText(),
			From: expr,
		}
	}

	return &mfast.Declare{
		Label:   ctx.Identifier().IDENTIFIER().GetText(),
		VarType: env.ARRAY,
		Size:    utils.ToInt(ctx.ArrayIndex().GetText()),
		Assign:  assign,
	}
}

func (v *AstGeneratorVisitor) VisitAssignment(ctx *mindfck.AssignmentContext) interface{} {
	expr := ctx.Expression().Accept(v).(mfast.Expr)

	var arrayIndex mfast.Expr
	if ctx.ArrayAccess() != nil {
		indexExpr := ctx.ArrayAccess().Expression().Accept(v).(mfast.Expr)
		arrayIndex = indexExpr
	}

	return &mfast.Assign{
		To:    ctx.Identifier().GetText(),
		From:  expr,
		Index: arrayIndex,
	}
}

func (v *AstGeneratorVisitor) VisitPrint(ctx *mindfck.PrintContext) interface{} {
	expr := ctx.Expression().Accept(v).(mfast.Expr)

	return &mfast.Print{
		Value: expr,
	}
}

func (v *AstGeneratorVisitor) VisitRead(ctx *mindfck.ReadContext) interface{} {
	return &mfast.Read{
		To: ctx.Identifier().GetText(),
	}
}

func (v *AstGeneratorVisitor) VisitDebug(ctx *mindfck.DebugContext) interface{} {
	return &mfast.Debug{}
}

func (v *AstGeneratorVisitor) VisitIfConditional(ctx *mindfck.IfConditionalContext) interface{} {
	condition := ctx.Expression().Accept(v).(mfast.Expr)
	block := ctx.Block(0).Accept(v).([]mfast.Stmt)
	elseBlock := []mfast.Stmt{}
	if ctx.Block(1) != nil {
		elseBlock = ctx.Block(1).Accept(v).([]mfast.Stmt)
	}

	return &mfast.If{
		Condition: condition,
		Block:     block,
		Else:      elseBlock,
	}
}

func (v *AstGeneratorVisitor) VisitWhileLoop(ctx *mindfck.WhileLoopContext) interface{} {
	return &mfast.While{
		Condition: ctx.Expression().Accept(v).(mfast.Expr),
		Block:     ctx.Block().Accept(v).([]mfast.Stmt),
	}
}
