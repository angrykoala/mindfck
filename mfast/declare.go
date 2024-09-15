package mfast

import (
	"mindfck/codegen"
)

type Declare struct {
	Label string
}

func (s *Declare) EvalStmt(cmd *codegen.CommandHandler) error {
	cmd.Declare(s.Label)
	// TODO: handle error of declare
	return nil
}
