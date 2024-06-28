package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + Literals
	IDENT TokenType = "IDENT" // foobar
	INT   TokenType = "INT"   // 12345

	// Operators
	ASSIGN  TokenType = "="
	ASTERIK TokenType = "*"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"

	// Data Query
	SELECT   TokenType = "SELECT"
	FROM     TokenType = "FROM"
	WHERE    TokenType = "WHERE"
	GROUPBY  TokenType = "GROUPBY"
	LIMIT    TokenType = "LIMIT"
	OFFSET   TokenType = "OFFSET"
	DISTINCT TokenType = "DISTINCT"

	// Data Maniplulation
	INSERT TokenType = "INSERT"
	INTO   TokenType = "INTO"
	VALUES TokenType = "VALUES"
	UPDATE TokenType = "UPDATE"
	SET    TokenType = "SET"
	DELETE TokenType = "DELETE"

	// Data Definition
	CREATE   TokenType = "CREATE"
	TABLE    TokenType = "TABLE"
	DATABASE TokenType = "DATABASE"
	INDEX    TokenType = "INDEX"
	VIEW     TokenType = "VIEW"
	DROP     TokenType = "DROP"
	ALTER    TokenType = "ALTER"
	TRUNCATE TokenType = "TRUNCATE"
)

type Token struct {
	Type    TokenType
	Literal string
}

func CreateToken(token TokenType, char byte) Token {
	return Token{
		Type:    token,
		Literal: string(char),
	}
}

var keyWords = map[string]TokenType{
	"SELECT": SELECT,
	"FROM":   FROM,
	"WHERE":  WHERE,
	"INSERT": INSERT,
	"DELETE": DELETE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keyWords[ident]; ok {
		return tok
	}
	return IDENT
}
