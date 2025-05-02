package tests

import (
	"testing"
	tk "github.com/RootControl/Interpreter/internal/token"
	lx "github.com/RootControl/Interpreter/internal/lexer"
)

func TestNextToken(t *testing.T) {
	input := "+={}(),;"

	tests := []struct {
		expectedType tk.TokenType
		expectedLiteral byte
	}{
		{tk.Plus, '+'},
		{tk.Assign, '='},
		{tk.LeftBrace, '{'},
		{tk.RightBrace, '}'},
		{tk.LeftParenthesis, '('},
		{tk.RightParenthesis, ')'},
		{tk.Comma, ','},
		{tk.Semicolon, ';'},
		{tk.EoF, 0},
	}

	lexer := lx.NewLexer(input)

	for i, tt := range tests {
		token, err := lexer.NextToken()

		if err != nil {
			t.Fatalf("tests[%d] - %s", i, err)
		}

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, token.Type)
		}

		if token.Literal != string(tt.expectedLiteral) {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
