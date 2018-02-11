package generator

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alanfoster/assembler/ast"
)

func TestAInstructionWithZero(t *testing.T) {
	g := New()
	result := g.ConvertBinary(&ast.AInstruction{
		Value: "0",
	})
	assert.Equal(t, "1000000000000000", result)
}

func TestAInstructionWithThree(t *testing.T) {
	g := New()
	result := g.ConvertBinary(&ast.AInstruction{
		Value: "3",
	})
	assert.Equal(t, "1000000000000011", result)
}

func TestAInstructionWithMaximumNumber(t *testing.T) {
	g := New()
	result := g.ConvertBinary(&ast.AInstruction{
		Value: "32767",
	})
	assert.Equal(t, "1111111111111111", result)
}
