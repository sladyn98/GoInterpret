package token

//TokenType is a string that specifies the type of token.
type TokenType string

//Token is a structure used to describe the token.
type Token struct {
	Type    TokenType
	Literal string
}

//Used to define the set of constants.
const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
)
