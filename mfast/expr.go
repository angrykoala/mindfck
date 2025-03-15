package mfast

import (
	"fmt"
	"mindfck/codegen"
	"mindfck/env"
)

type Expr interface {
	// Evaluates the expression, returning a single anonymous variable with the result, that must be released afterwards
	EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error)
}

type Literal struct {
	Value int
	Type  env.VarType
}

func (lit *Literal) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	res := cmd.Env().DeclareAnonVariable(lit.Type)

	switch lit.Type {
	case env.BYTE:
		cmd.SetByte(res, lit.Value)
	case env.INT:
		cmd.SetInt(res, lit.Value)

	}

	return res, nil
}

type ArrayLiteral struct {
	Value []int
	Type  env.VarType
}

func (lit *ArrayLiteral) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	res := cmd.Env().DeclareAnonArray(len(lit.Value))
	cmd.SetArray(res, lit.Value)

	return res, nil
}

type ArrayAccess struct {
	Target Expr
	Index  Expr
}

func (expr *ArrayAccess) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v1, err := expr.Target.EvalExpr(cmd)
	defer cmd.ReleaseIfAnonymous(v1)
	if err != nil {
		return nil, err
	}

	if v1.Type() != env.ARRAY {
		panic("Cannot access array, invalid type. Must be an array")
	}

	if w, ok := expr.Index.(*Literal); ok {
		res := cmd.Env().DeclareAnonByte()
		cmd.Copy(v1.GetByte(w.Value+env.ARRAY_HEAD_SIZE), res)
		return res, nil

	} else {
		index, err := expr.evaluateIndex(cmd)
		if err != nil {
			return nil, err
		}
		cmd.ReleaseIfAnonymous(index)

		res := cmd.Env().DeclareAnonByte()
		cmd.ReadIndex(v1, index, res)
		return res, nil
	}
}

func (expr *ArrayAccess) evaluateIndex(cmd *codegen.CommandHandler) (env.Variable, error) {
	index, err := expr.Index.EvalExpr(cmd)
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

type VariableExpr struct {
	Label string
}

func (lit *VariableExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v1 := cmd.Env().ResolveLabel(lit.Label)

	return v1, nil
}

type Operand string

const (
	PLUS     Operand = "+"
	MINUS    Operand = "-"
	MULTIPLY Operand = "*"
	DIVIDE   Operand = "/"

	// Booleans
	EQUALEQUAL Operand = "=="
	GT         Operand = ">"
	LT         Operand = "<"
	GTE        Operand = ">="
	LTE        Operand = "<="

	AND Operand = "and"
	OR  Operand = "or"
	NOT Operand = "not"
)

type BinaryExpr struct {
	Operator Operand
	Left     Expr
	Right    Expr
}

func (expr *BinaryExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v1, err := expr.Left.EvalExpr(cmd)
	if err != nil {
		return nil, err
	}
	defer cmd.ReleaseIfAnonymous(v1)

	v2, err := expr.Right.EvalExpr(cmd)
	if err != nil {
		return nil, err
	}
	defer cmd.ReleaseIfAnonymous(v2)

	if v1.Type() == env.BYTE && v2.Type() == env.INT {
		v3 := cmd.Env().DeclareAnonVariable(env.BYTE)
		defer cmd.Env().ReleaseVariable(v3)
		cmd.CastIntToByte(v2, v3)
		return expr.evalByteExpr(cmd, v1, v3)
	} else {
		codegen.AssertSameSize(v1, v2) //TODO: this should be in utils or something
		if v1.Type() == env.INT && v2.Type() == env.INT {
			return expr.evalIntExpr(cmd, v1, v2)
		} else {
			return expr.evalByteExpr(cmd, v1, v2)
		}
	}

}

func (expr *BinaryExpr) evalIntExpr(cmd *codegen.CommandHandler, v1 env.Variable, v2 env.Variable) (env.Variable, error) {
	switch expr.Operator {
	case PLUS:
		v3 := cmd.Env().DeclareAnonVariable(env.INT)
		cmd.AddInt(v1, v2, v3)
		return v3, nil
	case MINUS:
		v3 := cmd.Env().DeclareAnonVariable(env.INT)
		cmd.SubInt(v1, v2, v3)
		return v3, nil
	case MULTIPLY:
		v3 := cmd.Env().DeclareAnonVariable(env.INT)
		cmd.MultInt(v1, v2, v3)
		return v3, nil
	case DIVIDE:
		v3 := cmd.Env().DeclareAnonVariable(env.INT)
		cmd.DivInt(v1, v2, v3)
		return v3, nil
	case EQUALEQUAL:
		v3 := cmd.Env().DeclareAnonVariable(env.BYTE)
		cmd.EqualsInt(v1, v2, v3)
		return v3, nil
	case GT:
		v3 := cmd.Env().DeclareAnonVariable(env.BYTE)
		cmd.GtInt(v1, v2, v3)
		return v3, nil
	case LT:
		v3 := cmd.Env().DeclareAnonVariable(env.BYTE)
		cmd.GtInt(v2, v1, v3)
		return v3, nil
	case GTE:
		v3 := cmd.Env().DeclareAnonVariable(env.BYTE)
		cmd.GteInt(v1, v2, v3)
		return v3, nil
	case LTE:
		v3 := cmd.Env().DeclareAnonVariable(env.BYTE)
		cmd.GteInt(v2, v1, v3)
		return v3, nil
	default:
		return nil, fmt.Errorf("evalexpr: invalid int operator %v", expr.Operator)
	}
}

func (expr *BinaryExpr) evalByteExpr(cmd *codegen.CommandHandler, v1 env.Variable, v2 env.Variable) (env.Variable, error) {
	v3 := cmd.Env().DeclareAnonByte()

	switch expr.Operator {
	case PLUS:
		cmd.AddByte(v1, v2, v3)
	case MINUS:
		cmd.SubByte(v1, v2, v3)
	case MULTIPLY:
		cmd.MultByte(v1, v2, v3)
	case DIVIDE:
		cmd.DivByte(v1, v2, v3)
	case EQUALEQUAL:
		cmd.EqualsByte(v1, v2, v3)
	case GT:
		cmd.GtByte(v1, v2, v3)
	case LT:
		cmd.GtByte(v2, v1, v3)
	case GTE:
		cmd.GteByte(v1, v2, v3)
	case LTE:
		cmd.GteByte(v2, v1, v3)
	case AND: // Boolean OPS
		cmd.And(v1, v2, v3)
	case OR:
		cmd.Or(v1, v2, v3)

	default:
		return nil, fmt.Errorf("evalexpr: invalid byte operator %v", expr.Operator)
	}
	return v3, nil
}

type NotExpr struct {
	Expr Expr
}

func (n *NotExpr) EvalExpr(cmd *codegen.CommandHandler) (env.Variable, error) {
	v, err := n.Expr.EvalExpr(cmd)
	defer cmd.ReleaseIfAnonymous(v)
	if err != nil {
		return nil, err
	}
	res := cmd.Env().DeclareAnonByte()

	switch v.Type() {
	case env.BYTE:
		cmd.NotByte(v, res)
	case env.INT:
		cmd.NotInt(v, res)
	default:
		panic(fmt.Errorf("invalid type %s for not", v.Type()))
	}
	return res, nil
}
