package generator

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alanfoster/assembler/ast"
)

func TestAInstructionWithZero(t *testing.T) {
	g := New()
	instruction := &ast.AInstruction{
		Value: "0",
	}
	result := g.ConvertAInstruction(instruction)
	assert.Equal(t, "0000000000000000", result)
}

func TestAInstructionWithThree(t *testing.T) {
	g := New()
	instruction := &ast.AInstruction{
		Value: "3",
	}
	result := g.ConvertAInstruction(instruction)
	assert.Equal(t, "0000000000000011", result)
}

func TestAInstructionWithMaximumNumber(t *testing.T) {
	g := New()
	instruction := &ast.AInstruction{
		Value: "32767",
	}
	result := g.ConvertAInstruction(instruction)
	assert.Equal(t, "0111111111111111", result)
}

func TestCInstructionPrefixCommand(t *testing.T) {
	g := New()
	instruction := &ast.CInstruction{
		Destination: nil,
		Command:     ast.Command{Value: "!D"},
		Jump:        nil,
	}
	result := g.ConvertCInstruction(instruction)
	assert.Equal(t, "1110001101000000", result)
}

func TestCInstructionInfixCommand(t *testing.T) {
	g := New()
	instruction := &ast.CInstruction{
		Destination: nil,
		Command:     ast.Command{Value: "D+1"},
		Jump:        nil,
	}
	result := g.ConvertCInstruction(instruction)
	assert.Equal(t, "1110011111000000", result)
}

func TestCInstructionAssignment(t *testing.T) {
	g := New()
	instruction := &ast.CInstruction{
		Destination: &ast.Value{Value: "A"},
		Command:     ast.Command{Value: "D+1"},
		Jump:        nil,
	}
	result := g.ConvertCInstruction(instruction)
	assert.Equal(t, "1110011111100000", result)
}

func TestCInstructionAlwaysJump(t *testing.T) {
	g := New()
	instruction := &ast.CInstruction{
		Destination: nil,
		Command:     ast.Command{Value: "0"},
		Jump:        &ast.Value{Value: "JGT"},
	}
	result := g.ConvertCInstruction(instruction)
	assert.Equal(t, "1110101010000001", result)
}
