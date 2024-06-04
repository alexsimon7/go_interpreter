package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // an unknown token character
	EOF     = "EOF"     // singals end of file

	//Identifiers
	IDENT = "IDENT" // variable names
	INT   = "INT"   // integers

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Deliminators
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
