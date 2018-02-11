package assembler

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

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
