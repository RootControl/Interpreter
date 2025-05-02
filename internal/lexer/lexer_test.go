package lexer

import (
	"testing"
	tk "github.com/RootControl/Interpreter/internal/token"
	lx "github.com/RootControl/Interpreter/internal/lexer"
)

func TestNextToken(t *testing.T) {
	input := "+={}(),;"

	tests := []struct {
		expectedType tk.TokenType
		expectedLiteral uint32
	}{
		{tk.Plus, 21},
		{tk.Assign, 20},
		{tk.LeftBrace, 34},
		{tk.RightBrace, 35},
		{tk.LeftParenthesis, 32},
		{tk.RightParenthesis, 33},
		{tk.Comma, 30},
		{tk.Semicolon, 31},
		{tk.EoF, 1},
	}

	lexer := lx.NewLexer(input)

	for i, tt := range tests {
		token := lexer.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}

		if token.Literal != string(tt.expectedLiteral) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
