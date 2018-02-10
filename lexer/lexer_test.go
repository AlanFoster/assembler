package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alanfoster/assembler/token"
)

func TestArbitraryCharTokens(t *testing.T) {
	input := "@+-!;"
	l := New(input)
	expected := []token.Token{
		{Type: token.AT, Lexeme: "@"},
		{Type: token.OPERATOR, Lexeme: "+"},
		{Type: token.OPERATOR, Lexeme: "-"},
		{Type: token.OPERATOR, Lexeme: "!"},
		{Type: token.SEMICOLON, Lexeme: ";"},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}

func TestValue(t *testing.T) {
	input := "HELLOWORLD"
	l := New(input)
	expected := []token.Token{
		{Type: token.VALUE, Lexeme: "HELLOWORLD"},
		{Type: token.EOF, Lexeme: ""},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}

func TestValueThenChar(t *testing.T) {
	input := "HELLOWORLD;"
	l := New(input)
	expected := []token.Token{
		{Type: token.VALUE, Lexeme: "HELLOWORLD"},
		{Type: token.SEMICOLON, Lexeme: ";"},
		{Type: token.EOF, Lexeme: ""},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}

func TestAInstruction(t *testing.T) {
	input := "@1"
	l := New(input)
	expected := []token.Token{
		{Type: token.AT, Lexeme: "@"},
		{Type: token.VALUE, Lexeme: "1"},
		{Type: token.EOF, Lexeme: ""},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}

func TestCInstruction(t *testing.T) {
	input := "A=D+1;JGT"
	l := New(input)
	expected := []token.Token{
		{Type: token.VALUE, Lexeme: "A"},
		{Type: token.EQUALS, Lexeme: "="},
		{Type: token.VALUE, Lexeme: "D"},
		{Type: token.OPERATOR, Lexeme: "+"},
		{Type: token.VALUE, Lexeme: "1"},
		{Type: token.SEMICOLON, Lexeme: ";"},
		{Type: token.VALUE, Lexeme: "JGT"},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}
