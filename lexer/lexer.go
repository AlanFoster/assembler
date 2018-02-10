package lexer

import (
	"github.com/alanfoster/assembler/token"
	"bytes"
)

type Lexer struct {
	index   int
	current byte
	source  string
}

func New(source string) *Lexer {
	l := &Lexer{
		source: source,
	}

	l.next()
	return l
}

func (l *Lexer) Advance() token.Token {
	var tok token.Token

	switch l.current {
	case '@':
		tok = newCharToken(token.AT, l.current)
	case ';':
		tok = newCharToken(token.SEMICOLON, l.current)
	case '=':
		tok = newCharToken(token.EQUALS, l.current)
	case '+':
		tok = newCharToken(token.OPERATOR, l.current)
	case '!':
		tok = newCharToken(token.OPERATOR, l.current)
	case '-':
		tok = newCharToken(token.OPERATOR, l.current)
	case 0:
		tok = newStringToken(token.EOF, "")
	default:
		return newStringToken(token.VALUE, l.readValue())
	}

	l.next()
	return tok
}

func newCharToken(tokenType token.Type, ch byte) token.Token {
	return newStringToken(tokenType, string(ch))
}

func newStringToken(tokenType token.Type, s string) token.Token {
	return token.Token{Type: tokenType, Lexeme: s}
}

func (l *Lexer) next() {
	if l.index >= len(l.source) {
		l.current = 0 // Null byte
	} else {
		l.current = l.source[l.index]
	}

	l.index++
}

func (l *Lexer) readValue() string {
	var buf bytes.Buffer

	for l.isValue(l.current) {
		buf.WriteByte(l.current)
		l.next()
	}

	return buf.String()
}

func (l *Lexer) isValue(c byte) bool {
	return l.isDigit(c) || l.isLetter(c) || c == '_'
}

func (l *Lexer) isDigit(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z'
}

func (l *Lexer) isLetter(c byte) bool {
	return c >= '0' && c <= '9'
}
