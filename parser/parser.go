package parser

import (
	"fmt"
	"mindfck/parser/mfast"
	"mindfck/parser/tokens"
	"mindfck/utils"
)

func Parse(tokens []tokens.StmtTokens) ([]mfast.Stmt, error) {
	statements := []mfast.Stmt{}
	for _, stmtTokens := range tokens {
		stmt, err := parseStmt(stmtTokens)
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}

	return statements, nil
}

func parseStmt(stmtTokens []tokens.Token) (mfast.Stmt, error) {
	first, stmtTokens := stmtTokens[0], stmtTokens[1:]

	switch first.Kind {

	case tokens.BYTE:
		return parseDeclaration(stmtTokens)
	case tokens.IDENTIFIER:
		return parseAssignment(first, stmtTokens)

	}

	return nil, fmt.Errorf("unknown statement %v,  %v", first, stmtTokens)
}

func parseDeclaration(tk tokens.StmtTokens) (*mfast.Declare, error) {
	if len(tk) != 1 {
		return nil, fmt.Errorf("error in declaration %v", tk)
	}

	var labeltk = tk[0]

	if labeltk.Kind != tokens.IDENTIFIER {
		return nil, fmt.Errorf("error in declaration, invalid token %v", labeltk)
	}
	return &mfast.Declare{
		Label: labeltk.Txt,
	}, nil
}

func parseAssignment(identifier tokens.Token, tk []tokens.Token) (*mfast.Assign, error) {
	if len(tk) < 2 {
		return nil, fmt.Errorf("error in assignment %v", tk)
	}

	operator, exprTokens := tk[0], tk[1:]

	if operator.Kind != tokens.EQUALS {
		return nil, fmt.Errorf("error in declaration, invalid operator %s", operator.Txt)
	}

	fromExpr, err := parseExpr(exprTokens)
	if err != nil {
		return nil, err
	}

	return &mfast.Assign{
		To:   identifier.Txt,
		From: fromExpr,
	}, nil
}

func parseExpr(tk []tokens.Token) (mfast.Expr, error) {
	if len(tk) < 1 {
		return nil, fmt.Errorf("invalid expression %v", tk)
	}

	first, tk := tk[0], tk[1:]

	if first.Kind == tokens.IDENTIFIER || first.IsLiteral() {
		value := utils.ToInt(first.Txt)
		if len(tk) == 0 {
			return &mfast.Literal{
				Value: value,
			}, nil
		}

		// Binary
		operator, tk := tk[0], tk[1:]

		if !operator.IsBinaryOperator() {
			return nil, fmt.Errorf("invalid operator %v", operator)
		}

		rightExpr, err := parseExpr(tk)
		if err != nil {
			return nil, err
		}

		// TODO: handle complex operators in left
		leftExpr, err := parseExpr([]tokens.Token{first})
		if err != nil {
			return nil, err
		}

		return &mfast.BinaryExpr{
			Right:    rightExpr,
			Operator: operator.Kind,
			Left:     leftExpr,
		}, nil

	}

	return nil, fmt.Errorf("invalid expression %v %v", first, tk)
}
