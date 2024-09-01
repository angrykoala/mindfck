package tokens

type TokenKind int

const (
	UNKNOWN TokenKind = iota
	BYTE
	PRINT

	IDENTIFIER

	EQUALS

	// Operators
	PLUS
	MINUS
	MULTIPLY
	DIVIDE

	// Booleans
	EQUALEQUAL
	GT
	LT
	GTE
	LTE

	AND
	OR
	NOT

	// Literals
	NUMBER
)

type Token struct {
	Kind TokenKind
	Txt  string
}

func (tk *Token) IsBinaryOperator() bool {
	switch tk.Kind {
	case
		PLUS,
		MINUS,
		MULTIPLY,
		DIVIDE,
		EQUALEQUAL,
		AND,
		OR,
		GT,
		LT,
		GTE,
		LTE:
		return true
	}
	return false

}

func (tk *Token) IsLiteral() bool {
	return tk.Kind == NUMBER
}
