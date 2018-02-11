package generator

import (
	"github.com/alanfoster/assembler/ast"
	"strconv"
	"fmt"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) ConvertBinary(instruction ast.Node) string {
	switch instruction := instruction.(type) {
	case *ast.AInstruction:
		opCode := "1"
		value, err := strconv.ParseInt(instruction.Value, 10, 16)
		if err != nil {
			panic(err)
		}

		return fmt.Sprintf("%s%015b", opCode, value)
	default:
		panic("")
	}
}
