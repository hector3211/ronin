package ast

import (
	"sql-parser/internal/token"
	"testing"
)

func TestString(t *testing.T) {
	ast := &Ast{
		Statements: []Statement{
			&SelectStatement{
				Token:     token.Token{Type: token.SELECT, Literal: "SELECT"},
				Columns:   []string{"*"},
				TokenTwo:  token.Token{Type: token.FROM, Literal: "FROM"},
				TableName: "users",
			},
		},
	}

	if ast.String() != "SELECT * FROM users;" {
		t.Errorf("ast.String() failed, got: %q", ast.String())
	}
}
