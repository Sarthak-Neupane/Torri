package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 12345
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	MODULUS  = "%"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACKET = "["
	RBRACKET = "]"

	LT  = "<"
	GT  = ">"
	NOT = "!"

	NOTEQUAL = "!="
	EQUAL    = "=="

	// keywords
	FUNCTION = "fcn"
	LET       = "la"
	TRUE     = "true"
	FALSE    = "false"
	IF     = "yedi"
	ELSEIF     = "kita"
	ELSE    = "natra"
	RETURN   = "firta"
)

var keywords = map[string]TokenType{
	"fcn": FUNCTION,
	"la":  LET,
	"true": TRUE,
	"false": FALSE,
	"yedi": IF,
	"kita": IF,
	"natra": ELSE,
	"firta": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
