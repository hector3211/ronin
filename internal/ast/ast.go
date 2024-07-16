package ast

import (
	"bytes"
	"fmt"
	"sql-parser/internal/lexer"
	"sql-parser/internal/token"
)

type Node interface{}

type SelectStatement struct {
	Fields    []string // could be '*'
	TableName string
}

type InsertStatement struct {
	TableName string
	Columns   []string
	Values    []string
}

type Ast struct {
	l            *lexer.Lexer
	Statements   []token.Token
	position     int
	currentToken token.Token
	peekToken    token.Token
}

func NewAst(l *lexer.Lexer, tokens []token.Token) *Ast {
	return &Ast{
		l:          l,
		Statements: tokens,
		position:   0,
	}
}

func (a *Ast) NextToken() {
	a.currentToken = a.peekToken
	a.peekToken = a.l.NextToken()
}

func (a *Ast) Parse() (Node, error) {
	a.NextToken()

	switch a.currentToken.Type {
	case token.SELECT:
		return a.parseSelect()
	case token.INSERT:
		return a.parserInsert()
	default:
		return nil, fmt.Errorf("Unexpected token: %s", a.currentToken.Literal)
	}
}

func (a *Ast) parseSelect() (Node, error) {
	stmt := &SelectStatement{}
	a.NextToken()

	for a.currentToken.Type != token.SEMICOLON {
		if a.currentToken.Type == token.ASTERIK {
			stmt.Fields = append(stmt.Fields, a.currentToken.Literal)
			a.NextToken()
		}
		if a.currentToken.Type == token.LPAREN {
			curr := a.currentToken
			for curr.Type != token.RPAREN {
				if curr.Type != token.COMMA {
					stmt.Fields = append(stmt.Fields, curr.Literal)
				}
				a.NextToken()
			}
		}
		if a.currentToken.Type == token.IDENT {
			stmt.TableName = a.currentToken.Literal
		}
		a.NextToken()
	}

	return stmt, nil
}

func (a *Ast) parserInsert() (Node, error) {
	stmt := &InsertStatement{}
	a.NextToken()

	if a.currentToken.Type != token.INTO {
		return nil, fmt.Errorf("Expected INTO, got %s", a.currentToken.Literal)
	}
	a.NextToken()

	if a.currentToken.Type != token.IDENT {
		return nil, fmt.Errorf("Expected table name, got %s", a.currentToken.Literal)
	}
	stmt.TableName = a.currentToken.Literal + " "
	a.NextToken()

	if a.currentToken.Type != token.VALUES {
		return nil, fmt.Errorf("Expected VALUES, got %s", a.currentToken.Literal)
	}
	a.NextToken()

	if a.currentToken.Type != token.LPAREN {
		return nil, fmt.Errorf("Expected LPAREN (, got %s", a.currentToken.Literal)
	}
	a.NextToken()

	for a.currentToken.Type != token.RPAREN && a.currentToken.Type != token.EOF {
		if a.currentToken.Type == token.IDENT {
			stmt.Values = append(stmt.Values, a.currentToken.Literal)
		} else if a.currentToken.Type != token.COMMA {
			return nil, fmt.Errorf("Expected COMMA, got %s", a.currentToken.Literal)
		}
		a.NextToken()
	}

	if a.currentToken.Type != token.RPAREN {
		return nil, fmt.Errorf("Expected RPAREN ), got %s", a.currentToken.Literal)
	}

	return stmt, nil
}

func (a *Ast) String() string {
	var out bytes.Buffer
	stmt := a.Statements[0]

	switch stmt.Type {
	case token.SELECT:
		out.WriteString(stringifySelectSatement(a.Statements))
	case token.INSERT:
		out.WriteString(stringifyInsertSatement(a.Statements))
	}

	return out.String()
}

func stringifySelectSatement(tokens []token.Token) string {
	stmt := ""
	multipleSelects := false

	for _, tok := range tokens {
		stmt += tok.Literal
		if tok.Type == token.LPAREN {
			multipleSelects = true
		}

		if tok.Type == token.RPAREN {
			multipleSelects = false
		}

		if !multipleSelects && tok.Type != token.IDENT && tok.Type != token.SEMICOLON {
			stmt += " "
		}
	}
	return stmt
}

func stringifyInsertSatement(tokens []token.Token) string {
	stmt := ""
	multipleInserts := false

	for _, tok := range tokens {
		stmt += tok.Literal
		if tok.Type == token.LPAREN {
			multipleInserts = true
		}

		if !multipleInserts {
			stmt += " "
		}

	}
	return stmt
}
