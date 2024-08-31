package tokens

type TokenKind int

const (
	UNKNOWN TokenKind = iota
	BYTE
	IDENTIFIER

	EQUALS

	PLUS
	NUMBER
)

type Token struct {
	Kind TokenKind
	Txt  string
}

func (tk *Token) IsBinaryOperator() bool {
	return tk.Kind == PLUS
}

func (tk *Token) IsLiteral() bool {
	return tk.Kind == NUMBER
}
