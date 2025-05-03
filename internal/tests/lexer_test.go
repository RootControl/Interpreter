package tests

import (
	"testing"
	tk "github.com/RootControl/Interpreter/internal/token"
	lx "github.com/RootControl/Interpreter/internal/lexer"
)

func TestNextToken(t *testing.T) {
	input :=  `let five = 5;

	let ten = 10;

	let add = fn (x, y) {
		x + y;
	};

	let result = add(five, ten);

	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	`

	tests := []struct {
		expectedType tk.TokenType
		expectedLiteral string
	}{
		{tk.Let, "let"},
		{tk.Identifier, "five"},
		{tk.Assign, "="},
		{tk.Integer, "5"},
		{tk.Semicolon, ";"},
		{tk.Let, "let"},
		{tk.Identifier, "ten"},
		{tk.Assign, "="},
		{tk.Integer, "10"},
		{tk.Semicolon, ";"},
		{tk.Let, "let"},
		{tk.Identifier, "add"},
		{tk.Assign, "="},
		{tk.Function, "fn"},
		{tk.LeftParenthesis, "("},
		{tk.Identifier, "x"},
		{tk.Comma, ","},
		{tk.Identifier, "y"},
		{tk.RightParenthesis, ")"},
		{tk.LeftBrace, "{"},
		{tk.Identifier, "x"},
		{tk.Plus, "+"},
		{tk.Identifier, "y"},
		{tk.Semicolon, ";"},
		{tk.RightBrace, "}"},
		{tk.Semicolon, ";"},
		{tk.Let, "let"},
		{tk.Identifier, "result"},
		{tk.Assign, "="},		
		{tk.Identifier, "add"},
		{tk.LeftParenthesis, "("},
		{tk.Identifier, "five"},
		{tk.Comma, ","},
		{tk.Identifier, "ten"},
		{tk.RightParenthesis, ")"},
		{tk.Semicolon, ";"},
		{tk.Bang, "!"},
		{tk.Minus, "-"},
		{tk.Slash, "/"},
		{tk.Asterisk, "*"},
		{tk.Integer, "5"},
		{tk.Semicolon, ";"},
		{tk.Integer, "5"},
		{tk.LessThan, "<"},
		{tk.Integer, "10"},
		{tk.GreaterThan, ">"},
		{tk.Integer, "5"},
		{tk.Semicolon, ";"},
		{tk.If, "if"},
		{tk.LeftParenthesis, "("},
		{tk.Integer, "5"},
		{tk.LessThan, "<"},
		{tk.Integer, "10"},
		{tk.RightParenthesis, ")"},
		{tk.LeftBrace, "{"},
		{tk.Return, "return"},
		{tk.True, "true"},
		{tk.Semicolon, ";"},
		{tk.RightBrace, "}"},
		{tk.Else, "else"},
		{tk.LeftBrace, "{"},
		{tk.Return, "return"},
		{tk.False, "false"},
		{tk.Semicolon, ";"},
		{tk.RightBrace, "}"},
		{tk.EoF, string(byte(0))},
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
