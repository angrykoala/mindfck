package tokens

type TokenKind int

const (
	UNKNOWN TokenKind = iota
	BYTE
	IDENTIFIER

	EQUALS

	// Operators
	PLUS
	MINUS
	MULTIPLY
	DIVIDE

	// Booleans
	EQUALEQUAL
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
		OR:
		return true
	}
	return false

}

func (tk *Token) IsLiteral() bool {
	return tk.Kind == NUMBER
}
