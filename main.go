package main

import (
	"fmt"
	"github.com/alanfoster/assembler/assembler"
)

func main() {
	input := "@3"
	output := assembler.New().Convert(input)

	fmt.Println(output)
}
