package mfast

import (
	"fmt"
	"mindfck/codegen"
	"mindfck/env"
)

type Expr interface {
	// Evaluates the expression, returning a single anonymous variable with the result, that must be released afterwards
	EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error)
}

type Literal struct {
	Value int
}

func (lit *Literal) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	res := cmd.Env().DeclareAnonByte()
	cmd.SetByte(res, lit.Value)
	return res, nil
}

type VariableExpr struct {
	Label string
}

func (lit *VariableExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v1 := cmd.Env().ResolveLabel(lit.Label)
	v2 := cmd.Env().DeclareAnonByte()

	cmd.CopyByte(v1, v2)

	return v2, nil
}

type Operand string

const (
	PLUS     Operand = "+"
	MINUS    Operand = "-"
	MULTIPLY Operand = "*"
	DIVIDE   Operand = "/"

	// Booleans
	EQUALEQUAL Operand = "=="
	GT         Operand = ">"
	LT         Operand = "<"
	GTE        Operand = ">="
	LTE        Operand = "<="

	AND Operand = "and"
	OR  Operand = "or"
	NOT Operand = "not"
)

type BinaryExpr struct {
	Operator Operand
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

	v3 := cmd.Env().DeclareAnonByte()

	switch expr.Operator {
	case PLUS:
		cmd.Add(v1, v2, v3)
	case MINUS:
		cmd.Sub(v1, v2, v3)
	case MULTIPLY:
		cmd.Mult(v1, v2, v3)
	case DIVIDE:
		cmd.Div(v1, v2, v3)
	case EQUALEQUAL:
		cmd.Equals(v1, v2, v3)
	case AND:
		cmd.And(v1, v2, v3)
	case OR:
		cmd.Or(v1, v2, v3)
	case GT:
		cmd.Gt(v1, v2, v3)
	case LT:
		cmd.Gt(v2, v1, v3)
	case GTE:
		cmd.Gte(v1, v2, v3)
	case LTE:
		cmd.Gte(v2, v1, v3)

	default:
		return nil, fmt.Errorf("evalexpr: invalid operator %v", expr.Operator)
	}
	return v3, nil
}

type NotExpr struct {
	Expr Expr
}

func (n *NotExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v, err := n.Expr.EvalExpr(cmd)
	defer cmd.Release(v)
	if err != nil {
		return nil, err
	}

	res := cmd.Env().DeclareAnonByte()
	cmd.Not(v, res)
	return res, nil
}
