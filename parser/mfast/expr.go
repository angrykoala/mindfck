package mfast

import (
	"mindfck/codegen"
	"mindfck/parser/tokens"
)

type Expr interface {
	EvalExpr(cmd codegen.CommandHandler) string
}

type Literal struct {
	Value int
}

func (lit *Literal) EvalExpr(cmd codegen.CommandHandler) string {
	return ">"
}

type BinaryExpr struct {
	Operator tokens.TokenKind // Using TokenKind to simplify
	Left     Expr
	Right    Expr
}

func (lit *BinaryExpr) EvalExpr(cmd codegen.CommandHandler) string {
	return "caca"
}
