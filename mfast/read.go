package mfast

import "mindfck/codegen"

type Read struct {
	To string
}

func (s *Read) EvalStmt(cmd *codegen.CommandHandler) error {
	v := cmd.Env().ResolveLabel(s.To)
	cmd.Read(v)
	return nil
}
