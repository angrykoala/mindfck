package mfast

import (
	"mindfck/codegen"
	"mindfck/env"
)

type Declare struct {
	Label   string
	VarType env.VarType
	Size    int
	Assign  *Assign // Assignment in things like int x = 10
}

func (s *Declare) EvalStmt(cmd *codegen.CommandHandler) error {
	if s.VarType == env.ARRAY {
		cmd.DeclareArray(s.Label, s.Size)
	} else {
		cmd.Declare(s.Label, s.VarType)
	}

	if s.Assign != nil {
		s.Assign.EvalStmt(cmd)
	}
	// TODO: handle error of declare
	return nil
}
