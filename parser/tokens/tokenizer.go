package tokens

import (
	"fmt"
	"mindfck/utils"
	"strings"
)

type StmtTokens []*Token

func Tokenizer(code string) ([]StmtTokens, error) {

	stmtTxt := strings.FieldsFunc(code, func(c rune) bool {
		return c == '\n'
	})

	tokens := []StmtTokens{}
	for _, rawToken := range stmtTxt {
		stmt, err := tokenizeStmt(rawToken)
		if err != nil {
			return nil, err
		}
		if len(stmt) > 0 {
			tokens = append(tokens, stmt)
		}
	}

	return tokens, nil
}

func tokenizeStmt(rawStmt string) (StmtTokens, error) {
	tokensTxt := strings.FieldsFunc(strings.TrimSpace(rawStmt), func(c rune) bool {
		return c == ' '
	})
	tokens := []*Token{}
	for _, rawToken := range tokensTxt {
		kind := getTokenKind(rawToken)
		if kind == UNKNOWN {
			return nil, fmt.Errorf("Tokenizer: Invalid token %s", rawToken)
		}

		newToken := Token{
			Txt:  rawToken,
			Kind: kind,
		}
		tokens = append(tokens, &newToken)
	}

	return tokens, nil
}

func getTokenKind(tkn string) TokenKind {
	switch tkn {
	case "byte":
		return BYTE
	case "print":
		return PRINT
	case "=":
		return EQUALS
	case "+":
		return PLUS
	case "-":
		return MINUS
	case "*":
		return MULTIPLY
	case "/":
		return DIVIDE

		// Booleans
	case "==":
		return EQUALEQUAL
	case ">":
		return GT
	case "<":
		return LT
	case ">=":
		return GTE
	case "<=":
		return LTE
	case "and":
		return AND
	case "or":
		return OR
	case "not":
		return NOT
	default:
		if utils.IsInt(tkn) {
			return NUMBER
		}

		if isIdentifier(tkn) {
			return IDENTIFIER
		}
	}

	return UNKNOWN
}

func isIdentifier(str string) bool {
	firstChar := utils.GetFirstRune(str)
	return len(str) > 0 && utils.IsAlphabeticChar(firstChar) && utils.IsAlphaNumeric(str)
}
