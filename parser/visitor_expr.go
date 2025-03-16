package parser

import (
	"fmt"
	"mindfck/env"
	"mindfck/mfast"
	mindfck "mindfck/parser/antlr"
	"mindfck/utils"
)

func (v *AstGeneratorVisitor) VisitExpression(ctx *mindfck.ExpressionContext) interface{} {
	if ctx.Literal() != nil {
		return ctx.Literal().Accept(v)
	} else if ctx.Identifier() != nil {
		return &mfast.VariableExpr{
			Label: ctx.Identifier().GetText(),
		}
	} else if ctx.Expression(0) != nil && ctx.Expression(1) != nil {
		// Binary expr
		left := ctx.Expression(0).Accept(v).(mfast.Expr)
		right := ctx.Expression(1).Accept(v).(mfast.Expr)

		return &mfast.BinaryExpr{
			Operator: mfast.Operand(ctx.GetOp().GetText()),
			Left:     left,
			Right:    right,
		}
	} else if ctx.ArrayAccess() != nil {
		indexExpr := ctx.ArrayAccess().Expression().Accept(v).(mfast.Expr)
		arrayExpr := ctx.Expression(0).Accept(v).(mfast.Expr)

		return &mfast.ArrayAccess{
			Target: arrayExpr,
			Index:  indexExpr,
		}
	} else if ctx.Expression(0) != nil {
		// Unary expressions
		if ctx.NOT() != nil {
			return &mfast.NotExpr{
				Expr: ctx.Expression(0).Accept(v).(mfast.Expr),
			}
		}
		// Parenthesized expression
		return ctx.Expression(0).Accept(v).(mfast.Expr)
	}

	panic("Invalid expression in Visit expression")
}

func (v *AstGeneratorVisitor) VisitLiteral(ctx *mindfck.LiteralContext) interface{} {
	if ctx.NUMBER() != nil {
		return &mfast.Literal{
			Value: utils.ToInt(ctx.GetText()),
			Type:  env.INT,
		}
	} else if ctx.CHAR() != nil {
		return &mfast.Literal{
			Value: int(ctx.CHAR().GetText()[1]),
			Type:  env.BYTE,
		}
	} else if ctx.BYTE_NUMBER() != nil {
		txt := ctx.BYTE_NUMBER().GetText()
		return &mfast.Literal{
			Value: utils.ToInt(txt[:len(txt)-1]),
			Type:  env.BYTE,
		}
	} else if ctx.ArrayLiteral() != nil {
		return ctx.ArrayLiteral().Accept(v)

	} else {
		panic(fmt.Sprintf("invalid literal %s", ctx.GetText()))
	}
}

func (v *AstGeneratorVisitor) VisitArrayLiteral(ctx *mindfck.ArrayLiteralContext) interface{} {
	items := ctx.AllArrayItem()

	parsedItems := []int{}
	for _, s := range items {
		parsedArrayItem := utils.ToInt(s.GetText())

		parsedItems = append(parsedItems, parsedArrayItem)
	}
	return &mfast.ArrayLiteral{
		Value: parsedItems,
		Type:  env.ARRAY,
	}
}
