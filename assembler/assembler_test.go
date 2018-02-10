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
