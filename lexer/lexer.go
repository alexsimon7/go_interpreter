package lexer

import (
	"go_interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position - points to the current char
	readPosition int  // current reading position - points to the next char
	ch           byte // current char being read
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// The readChar method sets the ch field to the next character in the input string
// (pointed at by readPosition), sets the current position to the next position (readPosition)
// and increments readPosition by one to the next character in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// The readPosition method reads an identifier advancing the position field until a non-letter
// character is found and then returns the identifer
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// The readNumber method reads a Number advancing the position field until a non-number or
// decimal character is found and then returns the number or float
func (l *Lexer) readNumber() (string, token.TokenType) {
	tokenToUse := token.INT
	position := l.position

	for isDigit(l.ch) || l.ch == '.' && tokenToUse == "INT" {

		if l.ch == '.' && tokenToUse == "INT" {
			tokenToUse = token.FLOAT
		} else if l.ch == '.' && tokenToUse == "FLOAT" || tokenToUse == "ILLEGAL" {
			tokenToUse = token.ILLEGAL
		}

		l.readChar()
	}

	return l.input[position:l.position], token.TokenType(tokenToUse)

}

// The skipWhitespace method advances the field position past any spaces, tabs, newline,
// or carriage returns in the input string
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal, tok.Type = l.readNumber()

			// tok.Type = token.INT
			// tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// newToken take a tokenType and byte(ch) and returns a token with the tokenType as
// the Type and a string of ch as the Literal. Note: newToken does not accept strings
// as its second argument
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter returns the boolean true if the byte ch is a letter or an "_" and
// false if not.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && 'Z' <= ch || ch == '_'
}

// isDigit returns the boolean true if the byte ch is a digit or period and
// false if not
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}
