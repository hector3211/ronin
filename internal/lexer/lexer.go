package lexer

import "sql-parser/internal/token"

type Lexer struct {
	input        string
	position     int // Current position
	readPosition int // Next Position
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpaces()

	switch l.ch {
	case '=':
		tok = token.CreateToken(token.ASSIGN, '=')
	case ',':
		tok = token.CreateToken(token.COMMA, ',')
	case ';':
		tok = token.CreateToken(token.SEMICOLON, ';')
	case '(':
		tok = token.CreateToken(token.LPAREN, '(')
	case ')':
		tok = token.CreateToken(token.RPAREN, ')')
	case '*':
		tok = token.CreateToken(token.ASTERIK, '*')
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifer()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.CreateToken(token.ILLEGAL, l.ch)
		}
	}
	l.ReadChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.ReadChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readIdentifer() string {
	position := l.position
	for isLetter(l.ch) {
		l.ReadChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
