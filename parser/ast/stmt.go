package ast

type Stmt interface {
	EvalStmt()
}

type Declare struct {
	Label string
}

type Assign struct {
	To   string
	From Expr
}
