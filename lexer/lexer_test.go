package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	// input := `=+(){},;`

	// tests := []struct {
	// 	expectedType    token.TokenType
	// 	expectedLiteral string
	// }{
	// 	{token.ASSIGN, "="},
	// 	{token.PLUS, "+"},
	// 	{token.LPAREN, "("},
	// 	{token.RPAREN, ")"},
	// 	{token.LBRACE, "{"},
	// 	{token.RBRACE, "}"},
	// 	{token.COMMA, ","},
	// 	{token.SEMICOLON, ";"},
	// 	{token.EOF, ""},
	// }

	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "}"},
		{token.IDENT, "x"},
		// @TODO FINISH IMPLEMENTING TEST!
	}

	l := New(input) // New is creating a new Lexer struct and imputing the above tests into the input

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
	}
}
