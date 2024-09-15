package mfast

import (
	"mindfck/codegen"
	"mindfck/env"
)

type While struct {
	Condition Expr
	Block     []Stmt
}

func (s *While) EvalStmt(cmd *codegen.CommandHandler) error {
	v, err := s.Condition.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v)

	var v2 env.Variable
	var nestedError error
	cmd.While(v, func() {
		for _, stmt := range s.Block {
			nestedError := stmt.EvalStmt(cmd)
			if nestedError != nil {
				err = nestedError
				return
			}
		}

		v2, nestedError = s.Condition.EvalExpr(cmd)
		if nestedError != nil {
			err = nestedError
			return
		}
		cmd.Move(v2, v)
	})
	if err != nil {
		return err
	}

	cmd.Release(v2)
	return nil
}
