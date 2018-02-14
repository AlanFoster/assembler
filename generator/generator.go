package generator

import (
	"github.com/alanfoster/assembler/ast"
	"fmt"
	"github.com/alanfoster/assembler/symboltable"
)

const nullDestCode = "000"

var destCodes = map[string]string{
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

const nullJmpCode = "000"

var jmpCodes = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

var compCodes = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) ConvertAInstruction(instruction *ast.AInstruction, st symboltable.SymbolTable) string {
	var number int

	switch value := instruction.Value.(type) {
	case *ast.Number:
		number = value.Value
	case *ast.Variable:
		number = st[value.Name]
	default:
		panic(fmt.Errorf("unexpected value %v", value))
	}

	opCode := "0"
	return fmt.Sprintf("%s%015b", opCode, number)
}

func (g *Generator) ConvertCInstruction(instruction *ast.CInstruction) string {
	opCode := "111"
	compCode := g.compCode(instruction.Command)
	destCode := g.destCode(instruction.Destination)
	jmpCode := g.jmpCode(instruction.Jump)

	return fmt.Sprintf("%s%s%s%s", opCode, compCode, destCode, jmpCode)
}

func (g *Generator) compCode(command ast.Command) string {
	if value, ok := compCodes[command.Value]; ok {
		return value
	}

	panic("command not found")
}

func (g *Generator) destCode(dest *ast.Value) string {
	if dest == nil {
		return nullDestCode
	}

	if value, ok := destCodes[dest.Value]; ok {
		return value
	}

	panic("Unknown dest code")
}

func (g *Generator) jmpCode(jmp *ast.Value) string {
	if jmp == nil {
		return nullJmpCode
	}

	if value, ok := jmpCodes[jmp.Value]; ok {
		return value
	}

	panic("Unknown jmp code")
}
