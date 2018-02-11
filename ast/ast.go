package ast

import (
	"fmt"
	"bytes"
)

type Node interface {
	fmt.Stringer
}

type Instruction interface {
	Node
}

type Program struct {
	Instructions []Instruction
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, instruction := range p.Instructions {
		out.WriteString(instruction.String())
	}

	return out.String()
}

type AInstruction struct {
	Node
	Instruction

	Value string
}

func (a *AInstruction) String() string {
	return fmt.Sprintf("@%v", a.Value)
}

type Value struct {
	Node

	Value string
}

func (v *Value) String() string {
	return v.Value
}

// This representation does not accurately store its expression
// i.e. prefix, infix, or value - it's somewhat cheaty, to make lookup easier.
type Command struct {
	Node

	Value string
}

func (v *Command) String() string {
	return v.Value
}

type CInstruction struct {
	Node
	Instruction

	Destination *Value
	Command     Command
	Jump        *Value
}

func (c *CInstruction) String() string {
	var out bytes.Buffer

	if c.Destination != nil {
		out.WriteString(c.Destination.String())
		out.WriteString("=")
	}

	out.WriteString(c.Command.String())

	if c.Jump != nil {
		out.WriteString(";")
		out.WriteString(c.Jump.String())
	}

	return out.String()
}
