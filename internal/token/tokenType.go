package token

type TokenType uint32

const (
	Illigal TokenType = iota
	EoF = 1

	// Identifiers and literals
	Identifier = 10
	Integer = 11

	// Operators
	Assign = 20
	Plus = 21
	Minus = 22
	Asterisk = 23
	Slash = 24
	Equal = 25
	NotEqual = 26

	// Delimiters
	Comma = 30
	Semicolon = 31
	LeftParenthesis = 32
	RightParenthesis = 33
	LeftBrace = 34
	RightBrace = 35

	// Keywords
	Function = 40
	Let = 41
)

var keywords = map[string]TokenType{
	"fn": Function,
	"let": Let,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}
