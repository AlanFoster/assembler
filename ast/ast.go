package ast

import (
	"fmt"
	"bytes"
)

type Node interface {
	fmt.Stringer
}

type AInstruction struct {
	Node

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
