package parser

import (
	"mindfck/mfast"
	mindfck "mindfck/parser/antlr"

	"github.com/antlr4-go/antlr/v4"
)

type AstGeneratorVisitor struct {
	mindfck.BasemindfckVisitor
}

func (v *AstGeneratorVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// TOp level rule
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
