package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alanfoster/assembler/ast"
	"github.com/alanfoster/assembler/lexer"
)

func TestAInstruction(t *testing.T) {
	input := "@1337"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.AInstruction{
				Value: &ast.Number{Value: 1337},
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestCInstructionBasic(t *testing.T) {
	input := "A"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.CInstruction{
				Destination: nil,
				Command:     ast.Command{Value: "A"},
				Jump:        nil,
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestCInstructionPrefixCommand(t *testing.T) {
	input := "!D"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.CInstruction{
				Destination: nil,
				Command:     ast.Command{Value: "!D"},
				Jump:        nil,
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestCInstructionInfixCommand(t *testing.T) {
	input := "D+1"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.CInstruction{
				Destination: nil,
				Command:     ast.Command{Value: "D+1"},
				Jump:        nil,
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestCInstructionAssignment(t *testing.T) {
	input := "A=D+1"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.CInstruction{
				Destination: &ast.Value{Value: "A"},
				Command:     ast.Command{Value: "D+1"},
				Jump:        nil,
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestCInstructionAlwaysJump(t *testing.T) {
	input := "0;JGT"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.CInstruction{
				Destination: nil,
				Command:     ast.Command{Value: "0"},
				Jump:        &ast.Value{Value: "JGT"},
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestCInstructionAssignmentJump(t *testing.T) {
	input := "D+1;JGT"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.CInstruction{
				Destination: nil,
				Command:     ast.Command{Value: "D+1"},
				Jump:        &ast.Value{Value: "JGT"},
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestCInstructionMemoryCommand(t *testing.T) {
	input := "MD=M-1"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.CInstruction{
				Destination: &ast.Value{Value: "MD"},
				Command:     ast.Command{Value: "M-1"},
				Jump:        nil,
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}

func TestLabel(t *testing.T) {
	input := "(LOOP)"
	l := lexer.New(input)
	p := New(l)
	result := p.ParseProgram()
	expected := ast.Program{
		Instructions: []ast.Instruction{
			&ast.LInstruction{
				Value: "LOOP",
			},
		},
	}

	assert.Equal(t, expected, result)
	assert.Equal(t, input, result.String())
}
