package tokens

import (
	"fmt"
	"mindfck/utils"
	"strings"
)

type StmtTokens []Token

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

func getTokenKind(tkn string) TokenKind {
	if tkn == "byte" {
		return BYTE
	}

	if tkn == "=" {
		return EQUALS
	}
	if tkn == "+" {
		return PLUS
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
