package mfast

import "mindfck/codegen"

type Print struct {
	Value Expr
}

func (s *Print) EvalStmt(cmd *codegen.CommandHandler) error {
	v, err := s.Value.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v)
	cmd.Print(v)
	return nil
}
