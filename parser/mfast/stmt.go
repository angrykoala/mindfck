package mfast

type Stmt interface {
	EvalStmt()
}

type Declare struct {
	Label string
}

func (s *Declare) EvalStmt() {

}

type Assign struct {
	To   string
	From Expr
}

func (s *Assign) EvalStmt() {

}
