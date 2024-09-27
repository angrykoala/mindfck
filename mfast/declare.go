package mfast

import (
	"mindfck/codegen"
	"mindfck/env"
)

type Declare struct {
	Label   string
	VarType env.VarType
}

func (s *Declare) EvalStmt(cmd *codegen.CommandHandler) error {
	cmd.Declare(s.Label, s.VarType)
	// TODO: handle error of declare
	return nil
}
