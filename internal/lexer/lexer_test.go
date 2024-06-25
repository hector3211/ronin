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

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		t.Logf("getting [%q] expecting [%q]", tok.Type, tt.expectedType)

		if tok.Type != tt.expectedType {
			t.Fatalf("test [%d] expected %q got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Liteal != tt.expedtedLiteral {
			t.Fatalf("expected [%s] got [%s]", tt.expedtedLiteral, tok.Liteal)
		}
	}
}
