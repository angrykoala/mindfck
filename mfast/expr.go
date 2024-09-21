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
	Type  env.VarType
}

func (lit *Literal) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	res := cmd.Env().DeclareAnonVariable(lit.Type)

	switch lit.Type {
	case env.BYTE:
		cmd.SetByte(res, lit.Value)
	case env.INT:
		cmd.SetInt(res, lit.Value)

	}

	return res, nil
}

type VariableExpr struct {
	Label string
}

func (lit *VariableExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v1 := cmd.Env().ResolveLabel(lit.Label)
	v2 := cmd.Clone(v1)

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
	codegen.AssertSameSize(v1, v2) //TODO: this should be in utils or something
	if v1.Type() == env.INT && v2.Type() == env.INT {
		return expr.evalIntExpr(cmd, v1, v2)
	} else {
		return expr.evalByteExpr(cmd, v1, v2)
	}

}

func (expr *BinaryExpr) evalIntExpr(cmd *codegen.CommandHandler, v1 env.Variable, v2 env.Variable) (env.Variable, error) {
	v3 := cmd.Env().DeclareAnonVariable(env.INT)
	switch expr.Operator {
	case PLUS:
		cmd.AddInt(v1, v2, v3)
	case MINUS:
		cmd.SubInt(v1, v2, v3)
	case MULTIPLY:
		cmd.MultInt(v1, v2, v3)
	case DIVIDE:
		cmd.DivInt(v1, v2, v3)
	default:
		return nil, fmt.Errorf("evalexpr: invalid operator %v", expr.Operator)
	}
	return v3, nil
}

func (expr *BinaryExpr) evalByteExpr(cmd *codegen.CommandHandler, v1 env.Variable, v2 env.Variable) (env.Variable, error) {
	v3 := cmd.Env().DeclareAnonByte()

	switch expr.Operator {
	case PLUS:
		cmd.AddByte(v1, v2, v3)
	case MINUS:
		cmd.SubByte(v1, v2, v3)
	case MULTIPLY:
		cmd.MultByte(v1, v2, v3)
	case DIVIDE:
		cmd.DivByte(v1, v2, v3)
	case EQUALEQUAL:
		cmd.EqualsByte(v1, v2, v3)
	case GT:
		cmd.GtByte(v1, v2, v3)
	case LT:
		cmd.GtByte(v2, v1, v3)
	case GTE:
		cmd.GteByte(v1, v2, v3)
	case LTE:
		cmd.GteByte(v2, v1, v3)
	case AND: // Boolean OPS
		cmd.And(v1, v2, v3)
	case OR:
		cmd.Or(v1, v2, v3)

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
