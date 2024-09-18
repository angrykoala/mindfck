package mfast

import (
	"mindfck/codegen"
	"mindfck/env"
)

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

	// Implicit cast
	if v1.Type() == env.INT && v2.Type() == env.BYTE {
		cmd.CastByteToInt(v2, v1)
	} else if v1.Type() == env.BYTE && v2.Type() == env.INT {
		cmd.CastIntToByte(v2, v1)
	} else {
		cmd.MoveByte(v2, v1)
	}

	return nil
}
