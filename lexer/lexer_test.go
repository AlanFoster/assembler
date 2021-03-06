package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alanfoster/assembler/token"
)

func TestArbitraryCharTokens(t *testing.T) {
	input := "()@;=|&+-!"
	l := New(input)
	expected := []token.Token{
		{Type: token.LEFT_BRACKET, Lexeme: "("},
		{Type: token.RIGHT_BRACKET, Lexeme: ")"},
		{Type: token.AT, Lexeme: "@"},
		{Type: token.SEMICOLON, Lexeme: ";"},
		{Type: token.EQUALS, Lexeme: "="},
		{Type: token.OPERATOR, Lexeme: "|"},
		{Type: token.OPERATOR, Lexeme: "&"},
		{Type: token.OPERATOR, Lexeme: "+"},
		{Type: token.OPERATOR, Lexeme: "-"},
		{Type: token.OPERATOR, Lexeme: "!"},
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

func TestJump(t *testing.T) {
	input := `
		JGT
		JEQ
		JGE
		JLT
		JNE
		JLE
		JMP
	`
	l := New(input)
	expected := []token.Token{
		{Type: token.JUMP, Lexeme: "JGT"},
		{Type: token.JUMP, Lexeme: "JEQ"},
		{Type: token.JUMP, Lexeme: "JGE"},
		{Type: token.JUMP, Lexeme: "JLT"},
		{Type: token.JUMP, Lexeme: "JNE"},
		{Type: token.JUMP, Lexeme: "JLE"},
		{Type: token.JUMP, Lexeme: "JMP"},

		{Type: token.EOF, Lexeme: ""},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}

func TestAInstruction(t *testing.T) {
	input := "@1234 @Constant"
	l := New(input)
	expected := []token.Token{
		{Type: token.AT, Lexeme: "@"},
		{Type: token.NUMBER, Lexeme: "1234"},
		{Type: token.AT, Lexeme: "@"},
		{Type: token.VALUE, Lexeme: "Constant"},
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
		{Type: token.NUMBER, Lexeme: "1"},
		{Type: token.SEMICOLON, Lexeme: ";"},
		{Type: token.JUMP, Lexeme: "JGT"},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}

func TestCInstructionWithWhitespaceAndComments(t *testing.T) {
	input := `
		// Comment description
		A=D+1;JGT // Inline comment
		// Trailing comment/
		/`
	l := New(input)
	expected := []token.Token{
		{Type: token.VALUE, Lexeme: "A"},
		{Type: token.EQUALS, Lexeme: "="},
		{Type: token.VALUE, Lexeme: "D"},
		{Type: token.OPERATOR, Lexeme: "+"},
		{Type: token.NUMBER, Lexeme: "1"},
		{Type: token.SEMICOLON, Lexeme: ";"},
		{Type: token.JUMP, Lexeme: "JGT"},
		{Type: token.INVALID, Lexeme: "/"},
		{Type: token.EOF, Lexeme: ""},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}

func TestLabel(t *testing.T) {
	input := `
		($LABEL.FOO.BAR.BAZ)
	`
	l := New(input)
	expected := []token.Token{
		{Type: token.LEFT_BRACKET, Lexeme: "("},
		{Type: token.VALUE, Lexeme: "$LABEL.FOO.BAR.BAZ"},
		{Type: token.RIGHT_BRACKET, Lexeme: ")"},
		{Type: token.EOF, Lexeme: ""},
	}

	for _, expectedToken := range expected {
		assert.Equal(t, expectedToken, l.Advance())
	}
}
