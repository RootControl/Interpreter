package token

type Token struct {
	Type TokenType
	Literal string
}

func NewToken(t TokenType, l byte) Token {
	return Token{Type: t, Literal: string(l)}
}
