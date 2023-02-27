package lexer

import "github.com/pedro-git-projects/monkey-script/src/token"

// Lexer represents the actual language lexer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

// New takes a input string and returns a pointer to a Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar yields the next charater and advances the position
// in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // check for EOF
		l.char = 0
	} else {
		l.char = l.input[l.readPosition] // sets to next char
	}
	l.position = l.readPosition
	l.readPosition++ // read position is always one char ahead of position
}

// readNumber reads a char if is number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter checks if a charater is a letter
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// isDigit checks if a the character is a digit
func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

// readIdentifier reads a character and advances to the next letter
// until a non letter character is found.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// eatWhitespace strips the necessary characters for parsing
func (l *Lexer) eatWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

// newToken yields a token from a char
func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

// NextToken gets the character under examination and returns a token
// of that character. Before returning, however, it updates the position.
func (l *Lexer) NextToken() token.Token {
	tok := *new(token.Token)
	l.eatWhitespace()
	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}
	l.readChar()
	return tok
}
