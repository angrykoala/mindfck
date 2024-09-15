package mfast

import "mindfck/codegen"

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
