package lexer

import (
	tk "github.com/RootControl/Interpreter/internal/token"
)
type Lexer struct {
	Input string
	InputLenght uint32
	CurrentPosition uint32
	ReadPosition uint32
	CurrentChar byte
}

func NewLexer(input string) *Lexer {
	if input == "" {
		return nil
	}

	lexer := &Lexer{Input: input, InputLenght: uint32(len(input))}
	lexer.nextChar()

	return lexer 
}

func (l *Lexer) nextChar() {
	if l.ReadPosition >= l.InputLenght {
		l.CurrentChar = 0
	} else {
		l.CurrentChar = l.Input[l.ReadPosition]
	}

	l.CurrentPosition = l.ReadPosition
	l.ReadPosition++
}

func (l *Lexer) peekChar() byte {
	if l.ReadPosition >= l.InputLenght {
		return 0
	}

	return l.Input[l.ReadPosition]
}

func (l *Lexer) NextToken() (tk.Token, error) {
	var tokenType tk.TokenType
	identifier := ""

	l.skipWhitespace()

	switch l.CurrentChar {
	case '=':
		tokenType = tk.Assign

		if l.peekChar() == '=' {
			tokenType = tk.Equal
			identifier = string(l.CurrentChar)
			l.nextChar()
			identifier += string(l.CurrentChar)
		}
	case '+':
		tokenType = tk.Plus
	case '-':
		tokenType = tk.Minus
	case '*':
		tokenType = tk.Asterisk
	case '/':
		tokenType = tk.Slash
	case ';':
		tokenType = tk.Semicolon
	case ',':
		tokenType = tk.Comma
	case '(':
		tokenType = tk.LeftParenthesis
	case ')':
		tokenType = tk.RightParenthesis
	case '{':
		tokenType = tk.LeftBrace
	case '}':
		tokenType = tk.RightBrace
	case '<':
		tokenType = tk.LessThan
	case '>':
		tokenType = tk.GreaterThan
	case '!':
		tokenType = tk.Bang

		if l.peekChar() == '=' {
			tokenType = tk.NotEqual
			identifier = string(l.CurrentChar)
			l.nextChar()
			identifier += string(l.CurrentChar)
		}
	case 0:
		tokenType = tk.EoF
	default:
		tokenType, identifier = l.getIdentifier()
		return tk.NewToken(tokenType, identifier), nil
	}

	if identifier == "" {
		identifier = string(l.CurrentChar)
	}

	token := tk.NewToken(tokenType, identifier)
	l.nextChar()

	return token, nil
}

func (l *Lexer) skipWhitespace() {
	for l.CurrentChar == ' ' || l.CurrentChar == '\t' || 
			l.CurrentChar == '\n' || l.CurrentChar == '\r' {
				l.nextChar()
			}
}

func (l *Lexer) getIdentifier() (tokenType tk.TokenType, identifier string) {
	if isLetter(l.CurrentChar) {
		identifier = l.readIdentifier(isLetter)
		tokenType = tk.LookupIdent(identifier)

	} else if isNumber(l.CurrentChar) {
		identifier = l.readIdentifier(isNumber)
		tokenType = tk.Integer

	} else {
		identifier = ""
		tokenType = tk.Illigal
	}

	return tokenType, identifier
}

func (l *Lexer) readIdentifier(isValid func(byte) bool) string {
	initialPosition := l.CurrentPosition

	for isValid(l.CurrentChar) {
		l.nextChar()
	}

	return l.Input[initialPosition:l.CurrentPosition]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}
