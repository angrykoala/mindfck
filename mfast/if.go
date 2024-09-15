package mfast

import "mindfck/codegen"

type If struct {
	Condition Expr
	Block     []Stmt
	Else      []Stmt
}

func ProcessBlock(block *[]Stmt, cmd *codegen.CommandHandler) error {
	var err error = nil

	for _, stmt := range *block {
		nestedError := stmt.EvalStmt(cmd)
		if nestedError != nil {
			err = nestedError
			return err
		}
	}

	return err
}

func (s *If) EvalStmt(cmd *codegen.CommandHandler) error {
	v, err := s.Condition.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v)

	cmd.IfElse(v, func() { ProcessBlock(&s.Block, cmd) }, func() { ProcessBlock(&s.Else, cmd) })

	return err
}
