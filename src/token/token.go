package token

// List of language tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, x, y
	INT   = "INT"   // -1,0,1,2...

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// keywords is a map representing the type
// of the language Keywords.
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks if an identifier is keyword
// or a user defined ident and returns the corresponding type.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// TokenType represents the type of a token.
// ie: "integer", "right bracket".
type TokenType string

// Token represents a lexical analysis token.
// Type represents the token type and
// Literal represents the literal value of the token.
type Token struct {
	Type    TokenType
	Literal string
}
