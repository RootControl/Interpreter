package lexer

import (
	"errors"

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

func (l *Lexer) NextToken() (tk.Token, error) {
	var tokenType tk.TokenType

	switch l.CurrentChar {
	case '=':
		tokenType = tk.Assign
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
	case 0:
		tokenType = tk.EoF
	default:
		tokenType = tk.Identifier
	}

	if tokenType == tk.Illigal {
		return tk.Token{}, errors.New("unknown token")
	}

	token := tk.NewToken(tokenType, l.CurrentChar)
	l.nextChar()

	return token, nil
}
