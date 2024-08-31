package ast

import (
	"mindfck/codegen"
	"mindfck/parser"
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
	operator parser.TokenType // Using TokenType to simplify
	left     Expr
	right    Expr
}
