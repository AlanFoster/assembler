package assembler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func removeWhitespace(s string) string {
	lines := strings.Split(s, "\n")

	var trimmed []string
	for _, line := range lines {
		cleaned := strings.TrimSpace(line)
		if cleaned != "" {
			trimmed = append(trimmed, cleaned)
		}
	}
	return strings.Join(trimmed, "\n")
}

func TestSingleInstruction(t *testing.T) {
	input := "@3"
	result := New().Convert(input)
	expected := "0000000000000011"

	assert.Equal(t, expected, result)
}

func TestMultipleInstructions(t *testing.T) {
	input := "@3\nD=A\n@3"
	result := New().Convert(input)
	expected := "0000000000000011\n1110110000010000\n0000000000000011"

	assert.Equal(t, expected, result)
}

func TestMaxProgram(t *testing.T) {
	input := `
		@0
		D=M
		@1
		D=D-M
		@10
		D;JGT
		@1
		D=M
		@12
		0;JMP
		@0
		D=M
		@2
		M=D
		@14
		0;JMP
	`
	result := New().Convert(input)
	expected := removeWhitespace(`
		0000000000000000
		1111110000010000
		0000000000000001
		1111010011010000
		0000000000001010
		1110001100000001
		0000000000000001
		1111110000010000
		0000000000001100
		1110101010000111
		0000000000000000
		1111110000010000
		0000000000000010
		1110001100001000
		0000000000001110
		1110101010000111
	`)

	assert.Equal(t, expected, result)
}
