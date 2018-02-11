package assembler

import (
	"github.com/alanfoster/assembler/lexer"
	"github.com/alanfoster/assembler/parser"
	"github.com/alanfoster/assembler/generator"
	"strings"
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

	g := generator.New()

	var binary []string
	for _, instruction := range program.Instructions {
		binary = append(binary, g.ConvertBinary(instruction))
	}

	return strings.Join(binary, "\n")
}
