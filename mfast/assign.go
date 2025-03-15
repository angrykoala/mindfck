package mfast

import (
	"fmt"
	"mindfck/codegen"
	"mindfck/env"
)

type Assign struct {
	To    string
	From  Expr
	Index Expr
}

func (s *Assign) EvalStmt(cmd *codegen.CommandHandler) error {
	v1 := cmd.Env().ResolveLabel(s.To)
	v2, err := s.From.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.ReleaseIfAnonymous(v2)

	if s.Index != nil {
		// Array assigment

		if v1.Type() != env.ARRAY {
			return fmt.Errorf("cannot do index assign to non-array variable")
		}
		if v2.Type() == env.ARRAY {
			return fmt.Errorf("cannot assign array to array")
		}

		index, err := s.evaluateIndex(cmd)
		if err != nil {
			return err
		}

		if v2.Type() == env.INT {
			castValue := cmd.Env().DeclareAnonByte()
			defer cmd.Release(castValue)
			cmd.CastIntToByte(v2, castValue)
			cmd.WriteIndex(v1, index, castValue)
		} else {
			cmd.WriteIndex(v1, index, v2)
		}

	} else {
		// Normal assignment

		// Implicit cast
		if v1.Type() == env.INT && v2.Type() == env.BYTE {
			cmd.CastByteToInt(v2, v1)
		} else if v1.Type() == env.BYTE && v2.Type() == env.INT {
			cmd.CastIntToByte(v2, v1)
		} else {
			cmd.Copy(v2, v1)
		}
	}

	return nil
}

func (s *Assign) evaluateIndex(cmd *codegen.CommandHandler) (env.Variable, error) {
	index, err := s.Index.EvalExpr(cmd)
	if err != nil {
		return nil, err
	}

	if index.Type() == env.INT {
		indexByte := cmd.Env().DeclareAnonVariable(env.BYTE)
		cmd.CastIntToByte(index, indexByte)
		cmd.ReleaseIfAnonymous(index) // Release because the returned variable is not index
		return indexByte, nil
	}

	if index.Type() == env.BYTE {
		return index, nil
	} else {
		return nil, fmt.Errorf("Invalid index variable, must be byte or int")
	}

}
