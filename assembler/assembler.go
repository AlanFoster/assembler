package assembler

import (
	"github.com/alanfoster/assembler/lexer"
	"github.com/alanfoster/assembler/parser"
	"github.com/alanfoster/assembler/generator"
	"strings"
	"github.com/alanfoster/assembler/ast"
	"github.com/alanfoster/assembler/symboltable"
	"fmt"
	"strconv"
)

type Assembler struct {
}

func New() *Assembler {
	return &Assembler{}
}

func (a *Assembler) Convert(source string) string {
	l := lexer.New(source)
	p := parser.New(l)
	program := p.ParseProgram()

	st := a.buildSymbolTable(program)
	return a.generateBinary(program, st)
}

// Builds the symbol table containing labels and their corresponding
// ROM locations, as a first pass from the source file.
func (a *Assembler) buildSymbolTable(program ast.Program) *symboltable.SymbolTable {
	st := symboltable.New()

	// Track the ROM index. This will be incremented for each known instruction that
	// will be output to ROM
	romIndex := 0

	for _, instruction := range program.Instructions {
		switch instruction := instruction.(type) {
		case *ast.LInstruction:
			// Remember that labels do not get output to ROM
			st.Add(instruction.Value, romIndex)
		case *ast.AInstruction:
			romIndex++
		case *ast.CInstruction:
			romIndex++
		default:
			panic(fmt.Errorf("unexpected instruction %v", instruction))
		}
	}

	return st
}

// Generate the corresponding binary representation for a given program and symbol tree.
// In the first pass of the program, the symbol tree will been pre-populated with the
// pre-defined symbols, as well as user defined labels.
//
// In this second pass, we can now begin to generate the binary representation
func (a *Assembler) generateBinary(program ast.Program, st *symboltable.SymbolTable) string {
	g := generator.New()

	// A point to the next free memory slot for variable assignment
	// The first 15 slots are taken by 'Registers', therefore the next free slot is 16
	freeMemorySlotIndex := 16

	var binary []string
	for _, instruction := range program.Instructions {
		if _, ok := instruction.(*ast.LInstruction); ok {
			continue
		}

		switch instruction := instruction.(type) {
		case *ast.LInstruction:
			// Labels do not get output to ROM, they are pseudo instructions
		case *ast.AInstruction:
			// TODO: Unexpected mutation is unexpected
			if isVariable(instruction.Value) {
				newVariable := !st.Contains(instruction.Value)
				if newVariable {
					st.Add(instruction.Value, freeMemorySlotIndex)
					freeMemorySlotIndex++
				}

				address := st.Get(instruction.Value)
				instruction.Value = strconv.Itoa(address)
			}

			binary = append(binary, g.ConvertAInstruction(instruction))
		case *ast.CInstruction:
			binary = append(binary, g.ConvertCInstruction(instruction))
		default:
			panic(fmt.Errorf("unexpected instruction %v", instruction))
		}
	}

	return strings.Join(binary, "\n")
}

// TODO: This should be encoded within the AST
func isVariable(value string) bool {
	isDigit := value[0] >= '0' && value[0] <= '9'
	return !isDigit
}
