package token

type Token struct {
	Type TokenType
	Literal string
}

func NewToken(t TokenType, l string) Token {
	return Token{Type: t, Literal: l}
}
