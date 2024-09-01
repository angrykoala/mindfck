package mfast

import (
	"fmt"
	"mindfck/codegen"
	"mindfck/env"
	"mindfck/parser/tokens"
)

type Expr interface {
	EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error)
}

type Literal struct {
	Value int
}

func (lit *Literal) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	res := cmd.Env().DeclareAnonVariable()
	cmd.Set(res, lit.Value)
	return res, nil
}

type VariableExpr struct {
	Label string
}

func (lit *VariableExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v1 := cmd.Env().ResolveLabel(lit.Label)
	v2 := cmd.Env().DeclareAnonVariable()

	cmd.Copy(v1, v2)

	return v2, nil
}

type BinaryExpr struct {
	Operator tokens.TokenKind // Using TokenKind to simplify
	Left     Expr
	Right    Expr
}

func (expr *BinaryExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v1, err := expr.Left.EvalExpr(cmd)
	if err != nil {
		return nil, err
	}
	defer cmd.Release(v1)
	v2, err := expr.Right.EvalExpr(cmd)
	if err != nil {
		return nil, err
	}
	defer cmd.Release(v2)

	v3 := cmd.Env().DeclareAnonVariable()

	switch expr.Operator {
	case tokens.PLUS:
		cmd.Add(v1, v2, v3)
	case tokens.MINUS:
		cmd.Sub(v1, v2, v3)
	case tokens.MULTIPLY:
		cmd.Mult(v1, v2, v3)
	case tokens.DIVIDE:
		cmd.Div(v1, v2, v3)
	case tokens.EQUALEQUAL:
		cmd.Equals(v1, v2, v3)
	case tokens.AND:
		cmd.And(v1, v2, v3)
	case tokens.OR:
		cmd.Or(v1, v2, v3)

	case tokens.GT:
		cmd.Gt(v1, v2, v3)
	case tokens.LT:
		cmd.Gte(v2, v1, v3)
	case tokens.GTE:
		cmd.Gte(v1, v2, v3)
	case tokens.LTE:
		cmd.Gt(v2, v1, v3)

	default:
		return nil, fmt.Errorf("evalexpr: invalid operator %v", expr.Operator)
	}
	return v3, nil
}
