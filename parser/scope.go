package parser

import "mindfck/parser/mfast"

type Scope struct {
	stmts []mfast.Stmt
	exprs []mfast.Expr
}

func (s *Scope) clearExpr() {
	s.exprs = []mfast.Expr{}
}

func (s *Scope) pushExpr(expr mfast.Expr) {
	s.exprs = append(s.exprs, expr)
}

func (s *Scope) popLeftExpr() mfast.Expr {
	if len(s.exprs) == 0 {
		panic("Can't popLeftExpr, expr is empty")
	}
	res := s.exprs[0]
	s.exprs = s.exprs[1:len(s.exprs)]

	return res
}

func (s *Scope) appendStmt(stmt mfast.Stmt) {
	s.stmts = append(s.stmts, stmt)
}

func newScope() Scope {
	return Scope{
		stmts: []mfast.Stmt{},
		exprs: []mfast.Expr{},
	}
}
