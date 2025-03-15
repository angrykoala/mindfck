package mfast

import (
	"mindfck/codegen"
)

type Debug struct {
}

func (s *Debug) EvalStmt(cmd *codegen.CommandHandler) error {
	cmd.DebugBreak()
	return nil
}
