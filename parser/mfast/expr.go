package mfast

import (
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

	cmd.Add(v1, v2, v3)
	return v3, nil
}
