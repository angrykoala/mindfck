package parser

import (
	"fmt"
	"mindfck/utils"
	"strings"
)

type TokenType int

const (
	UNKNOWN TokenType = iota
	BYTE
	EQUALS
	IDENTIFIER
	NUMBER
)

type Token struct {
	Kind TokenType
	Txt  string
}

type stmt []Token

func Tokenizer(rawStmt string) (stmt, error) {

	tokensTxt := strings.FieldsFunc(rawStmt, func(c rune) bool {
		return c == ' '
	})

	tokens := []Token{}
	for _, rawToken := range tokensTxt {
		kind := getTokenKind(rawToken)
		if kind == UNKNOWN {
			return nil, fmt.Errorf("Tokenizer: Invalid token %s", rawToken)
		}

		newToken := Token{
			Txt:  rawToken,
			Kind: kind,
		}
		tokens = append(tokens, newToken)
	}

	return tokens, nil
}

func getTokenKind(tkn string) TokenType {
	if tkn == "byte" {
		return BYTE
	}

	if tkn == "=" {
		return EQUALS
	}

	if utils.IsInt(tkn) {
		return NUMBER
	}

	if isIdentifier(tkn) {
		return IDENTIFIER
	}

	return UNKNOWN
}

func isIdentifier(str string) bool {
	firstChar := utils.GetFirstRune(str)
	return len(str) > 0 && utils.IsAlphabeticChar(firstChar) && utils.IsAlphaNumeric(str)
}
