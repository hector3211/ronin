package ast

import (
	"sql-parser/internal/token"
	"testing"
)

func TestString(t *testing.T) {
	tokens := []token.Token{
		{
			Type:    token.SELECT,
			Literal: "SELECT",
		},
		{
			Type:    token.ASTERIK,
			Literal: "*",
		},
		{
			Type:    token.FROM,
			Literal: "FROM",
		},
		{
			Type:    token.IDENT,
			Literal: "users",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
	}

	ast := NewAst(tokens)

	if ast.String() != "SELECT * FROM users;" {
		t.Errorf("ast.String() failed, got: %q", ast.String())
	}
}

func TestStringTwo(t *testing.T) {
	tokens := []token.Token{
		{
			Type:    token.INSERT,
			Literal: "INSERT",
		},
		{
			Type:    token.INTO,
			Literal: "INTO",
		},
		{
			Type:    token.IDENT,
			Literal: "users",
		},
		{
			Type:    token.VALUES,
			Literal: "VALUES",
		},

		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "maddog",
		},
		{
			Type:    token.COMMA,
			Literal: ",",
		},
		{
			Type:    token.IDENT,
			Literal: "dogmad",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
	}

	ast := NewAst(tokens)

	if ast.String() != "INSERT INTO users VALUES (maddog,dogmad);" {
		t.Errorf("ast.String() failed, got: %q", ast.String())
	}
}
