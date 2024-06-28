package lexer

import (
	"sql-parser/internal/token"
	"testing"
)

func TestOne(t *testing.T) {
	input := `SELECT * FROM user;`

	tests := []struct {
		expectedType    token.TokenType
		expedtedLiteral string
	}{
		{token.SELECT, "SELECT"},
		{token.ASTERIK, "*"},
		{token.FROM, "FROM"},
		{token.IDENT, "user"},
		{token.SEMICOLON, ";"},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()

		t.Logf("getting [%q] expecting [%q]", tok.Type, tt.expectedType)

		if tok.Type != tt.expectedType {
			t.Fatalf("test [%d] expected %q got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expedtedLiteral {
			t.Fatalf("expected [%s] got [%s]", tt.expedtedLiteral, tok.Literal)
		}
	}
}
