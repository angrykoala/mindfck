package parser

import (
	"mindfck/mfast"
	mindfck "mindfck/parser/antlr"
	"mindfck/utils"

	"github.com/antlr4-go/antlr/v4"
)

type AstGeneratorVisitor struct {
	mindfck.BasemindfckVisitor
}

func (v *AstGeneratorVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *AstGeneratorVisitor) VisitStatements(ctx *mindfck.StatementsContext) interface{} {
	result := []mfast.Stmt{}
	if ctx.AllStatement() != nil {
		for _, s := range ctx.AllStatement() {
			stmt := s.Accept(v).(mfast.Stmt)

			result = append(result, stmt)
		}
	}

	return result
}

func (v *AstGeneratorVisitor) VisitStatement(ctx *mindfck.StatementContext) interface{} {
	if ctx.Declaration() != nil {
		return ctx.Declaration().Accept(v)
	}
	if ctx.Declaration() != nil {
		return ctx.Declaration().Accept(v)
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

	panic("Spanish Inquisition")
}

func (v *AstGeneratorVisitor) VisitDeclaration(ctx *mindfck.DeclarationContext) interface{} {
	return &mfast.Declare{
		Label: ctx.Identifier().IDENTIFIER().GetText(),
	}
}

func (v *AstGeneratorVisitor) VisitAssignment(ctx *mindfck.AssignmentContext) interface{} {
	expr := ctx.Expression().Accept(v).(mfast.Expr)

	return &mfast.Assign{
		To:   ctx.Identifier().GetText(),
		From: expr,
	}
}
func (v *AstGeneratorVisitor) VisitPrint(ctx *mindfck.PrintContext) interface{} {
	expr := ctx.Expression().Accept(v).(mfast.Expr)

	return &mfast.Print{
		Value: expr,
	}
}

func (v *AstGeneratorVisitor) VisitExpression(ctx *mindfck.ExpressionContext) interface{} {
	if ctx.Literal() != nil {
		return &mfast.Literal{
			Value: utils.ToInt(ctx.Literal().GetText()),
		}

	} else if ctx.Identifier() != nil {
		return &mfast.VariableExpr{
			Label: ctx.Identifier().GetText(),
		}
	} else if ctx.Expression(0) != nil && ctx.Expression(1) != nil {
		left := ctx.Expression(0).Accept(v).(mfast.Expr)
		right := ctx.Expression(1).Accept(v).(mfast.Expr)

		return &mfast.BinaryExpr{
			Operator: mfast.Operand(ctx.GetOp().GetText()),
			Left:     left,
			Right:    right,
		}
	} else if ctx.Expression(0) != nil {
		if ctx.NOT() != nil {
			return &mfast.NotExpr{
				Expr: ctx.Expression(0).Accept(v).(mfast.Expr),
			}
		}

		return ctx.Expression(0).Accept(v).(mfast.Expr)
	}

	panic("Invalid expression in Visit expression")
}

func (v *AstGeneratorVisitor) VisitIfConditional(ctx *mindfck.IfConditionalContext) interface{} {
	return &mfast.If{
		Condition: ctx.Expression().Accept(v).(mfast.Expr),
		Block:     ctx.Block().Accept(v).([]mfast.Stmt),
	}
}

func (v *AstGeneratorVisitor) VisitBlock(ctx *mindfck.BlockContext) interface{} {
	result := []mfast.Stmt{}
	if ctx.AllStatement() != nil {
		for _, s := range ctx.AllStatement() {
			stmt := s.Accept(v).(mfast.Stmt)

			result = append(result, stmt)
		}
	}

	return result
}

func (v *AstGeneratorVisitor) VisitWhileLoop(ctx *mindfck.WhileLoopContext) interface{} {
	return &mfast.While{
		Condition: ctx.Expression().Accept(v).(mfast.Expr),
		Block:     ctx.Block().Accept(v).([]mfast.Stmt),
	}
}
