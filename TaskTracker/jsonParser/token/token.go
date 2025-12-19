package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	STRING = "STRING"
	NUMBER = "NUMBER"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"

	// Delimiters
	COMMA = ","
	COLON = ":"

	// Brackets
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
)
