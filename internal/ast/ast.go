package ast

import (
	"bytes"
	"sql-parser/internal/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type SelectStatement struct {
	Token     token.Token
	Columns   []string // could be '*'
	TokenTwo  token.Token
	TableName string
}

func (s *SelectStatement) statementNode() {}
func (s *SelectStatement) TokenLiteral() string {
	return s.Token.Literal
}

func (s *SelectStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.TokenLiteral() + " ")
	for _, col := range s.Columns {
		out.WriteString(col + " ")
	}

	out.WriteString(s.TokenTwo.Literal + " ")
	out.WriteString(s.TableName)
	out.WriteString(";")

	return out.String()
}

type Ast struct {
	Statements []Statement
}

func (a *Ast) String() string {
	var out bytes.Buffer

	for _, s := range a.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (a *Ast) TokenLiteral() string {
	if len(a.Statements) > 0 {
		return a.Statements[0].TokenLiteral()
	}

	return ""
}
