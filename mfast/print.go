package mfast

import (
	"fmt"
	"mindfck/codegen"
	"mindfck/env"
)

type Print struct {
	Value Expr
}

func (s *Print) EvalStmt(cmd *codegen.CommandHandler) error {
	v, err := s.Value.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v)
	switch v.Type() {
	case env.BYTE:
		cmd.PrintByte(v)
	case env.INT:
		cmd.PrintInt(v)
	default:
		panic(fmt.Errorf("cannot print variable of type %s", v.Type()))
	}
	return nil
}
