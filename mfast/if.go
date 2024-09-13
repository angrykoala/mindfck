package mfast

import "mindfck/codegen"

type If struct {
	Condition Expr
	Block     []Stmt
}

func (s *If) EvalStmt(cmd *codegen.CommandHandler) error {
	v, err := s.Condition.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v)

	cmd.If(v, func() {
		for _, stmt := range s.Block {
			nestedError := stmt.EvalStmt(cmd)
			if nestedError != nil {
				err = nestedError
				return
			}
		}
	})

	return err
}
