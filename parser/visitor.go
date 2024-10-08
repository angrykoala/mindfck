package parser

import (
	"fmt"
	"mindfck/env"
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
	if ctx.Read() != nil {
		return ctx.Read().Accept(v)
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

	return &mfast.Declare{
		Label:   ctx.Identifier().IDENTIFIER().GetText(),
		VarType: varType,
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

func (v *AstGeneratorVisitor) VisitRead(ctx *mindfck.ReadContext) interface{} {
	return &mfast.Read{
		To: ctx.Identifier().GetText(),
	}
}

func (v *AstGeneratorVisitor) VisitExpression(ctx *mindfck.ExpressionContext) interface{} {
	if ctx.Literal() != nil {
		if ctx.Literal().NUMBER() != nil {
			return &mfast.Literal{
				Value: utils.ToInt(ctx.Literal().GetText()),
				Type:  env.INT,
			}
		} else if ctx.Literal().CHAR() != nil {
			return &mfast.Literal{
				Value: int(ctx.Literal().CHAR().GetText()[1]),
				Type:  env.BYTE,
			}
		} else if ctx.Literal().BYTE_NUMBER() != nil {
			txt := ctx.Literal().BYTE_NUMBER().GetText()
			return &mfast.Literal{
				Value: utils.ToInt(txt[:len(txt)-1]),
				Type:  env.BYTE,
			}
		} else {
			panic(fmt.Sprintf("invalid literal %s", ctx.Literal().GetText()))
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
