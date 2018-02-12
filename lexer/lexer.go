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

	l.skipCommentsAndWhitespace()

	switch l.current {
	case '(':
		tok = newCharToken(token.LEFT_BRACKET, l.current)
	case ')':
		tok = newCharToken(token.RIGHT_BRACKET, l.current)
	case '@':
		tok = newCharToken(token.AT, l.current)
	case ';':
		tok = newCharToken(token.SEMICOLON, l.current)
	case '=':
		tok = newCharToken(token.EQUALS, l.current)
	case '|':
		tok = newCharToken(token.OPERATOR, l.current)
	case '&':
		tok = newCharToken(token.OPERATOR, l.current)
	case '+':
		tok = newCharToken(token.OPERATOR, l.current)
	case '-':
		tok = newCharToken(token.OPERATOR, l.current)
	case '!':
		tok = newCharToken(token.OPERATOR, l.current)
	case 0:
		tok = newStringToken(token.EOF, "")
	default:
		if l.isValue(l.current) {
			value := l.readValue()
			tokenType := token.MapValue(value)

			return newStringToken(tokenType, value)
		} else {
			tok = newCharToken(token.INVALID, l.current)
		}
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

func (l *Lexer) skipCommentsAndWhitespace() {
	for {
		l.skipWhitespace()
		hasSkippedComments := l.skipComments()
		if !hasSkippedComments {
			break
		}
	}
}

func (l *Lexer) skipComments() bool {
	if l.isComment() {
		for l.current != '\n' && l.current != 0 {
			l.next()
		}
		return true
	}
	return false
}

func (l *Lexer) isComment() bool {
	return l.current == '/' && l.peek() == '/'
}

func (l *Lexer) peek() byte {
	return l.source[l.index]
}

func (l *Lexer) skipWhitespace() {
	for l.isWhitespace(l.current) {
		l.next()
	}
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
	return l.isDigit(c) || l.isLetter(c) ||  c == '.' || c == '_' || c == '$'
}

func (l *Lexer) isDigit(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z'
}

func (l *Lexer) isLetter(c byte) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) isWhitespace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t'
}
