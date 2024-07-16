package ast

import (
	"sql-parser/internal/lexer"
	"sql-parser/internal/token"
	"testing"
)

func TestOne(t *testing.T) {
	input := "SELECT * FROM users;"
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

	l := lexer.New(input)
	ast := NewAst(l, tokens)

	if ast.String() != input {
		t.Errorf("ast.String() failed, got: %q", ast.String())
	}
}

func TestTwo(t *testing.T) {
	input := "SELECT (name,age) FROM users;"
	tokens := []token.Token{
		{
			Type:    token.SELECT,
			Literal: "SELECT",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.IDENT,
			Literal: "name",
		},
		{
			Type:    token.COMMA,
			Literal: ",",
		},
		{
			Type:    token.IDENT,
			Literal: "age",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
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

	l := lexer.New(input)
	ast := NewAst(l, tokens)

	if ast.String() != input {
		t.Errorf("ast.String() failed, got: %q", ast.String())
	}
}

func TestThree(t *testing.T) {
	input := "INSERT INTO users VALUES (maddog,dogmad);"
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

	l := lexer.New(input)
	ast := NewAst(l, tokens)

	if ast.String() != input {
		t.Errorf("ast.String() failed, got: %q", ast.String())
	}
}
