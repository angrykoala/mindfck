package mfast

import (
	"mindfck/codegen"
	"mindfck/env"
)

type Stmt interface {
	EvalStmt(cmd *codegen.CommandHandler) error
}

type Declare struct {
	Label string
}

func (s *Declare) EvalStmt(cmd *codegen.CommandHandler) error {
	cmd.Declare(s.Label)
	// TODO: handle error of declare
	return nil
}

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

type If struct {
	Condition Expr
	Block     []Stmt
}

func (s *If) EvalStmt(cmd *codegen.CommandHandler) error {
	v, err := s.Condition.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v)

	cmd.If(v, func() {
		for _, stmt := range s.Block {
			nestedError := stmt.EvalStmt(cmd)
			if nestedError != nil {
				err = nestedError
				return
			}
		}
	})

	return err
}

type While struct {
	Condition Expr
	Block     []Stmt
}

func (s *While) EvalStmt(cmd *codegen.CommandHandler) error {
	v, err := s.Condition.EvalExpr(cmd)
	if err != nil {
		return err
	}
	defer cmd.Release(v)

	var v2 env.Variable
	var nestedError error
	cmd.While(v, func() {
		for _, stmt := range s.Block {
			nestedError := stmt.EvalStmt(cmd)
			if nestedError != nil {
				err = nestedError
				return
			}
		}

		v2, nestedError = s.Condition.EvalExpr(cmd)
		if nestedError != nil {
			err = nestedError
			return
		}
		cmd.Move(v2, v)
	})
	if err != nil {
		return err
	}

	cmd.Release(v2)
	return nil
}
