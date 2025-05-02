package token

type TokenType uint32

type Token struct {
	Type TokenType
	Literal string
}
