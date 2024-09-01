package mfast

import "mindfck/codegen"

type Stmt interface {
	EvalStmt(cmd *codegen.CommandHandler) error
}

type Declare struct {
	Label string
}

func (s *Declare) EvalStmt(cmd *codegen.CommandHandler) error {
	cmd.Declare(s.Label)
	// TODO: handle error of declare
	return nil
}

type Assign struct {
	To   string
	From Expr
}

func (s *Assign) EvalStmt(cmd *codegen.CommandHandler) error {
	v1 := cmd.Env().ResolveLabel(s.To)

	v2, err := s.From.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v2)
	cmd.Move(v2, v1)

	return nil
}
