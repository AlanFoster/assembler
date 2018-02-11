package main

import (
	"github.com/alanfoster/assembler/assembler"
	"flag"
	"io/ioutil"
	"fmt"
)

func assemble(entryFile string, outputFile string) {
	data, err := ioutil.ReadFile(entryFile)
	if err != nil {
		fmt.Println("Ruh roh")
		panic(err)
	}
	source := string(data)
	result := assembler.New().Convert(source)

	ioutil.WriteFile(outputFile, []byte(result), 0644)
}

func main() {
	var entryFile string
	var outputFile string
	flag.StringVar(&entryFile, "entry-file", "", "File to convert to hack")
	flag.StringVar(&outputFile, "output-file", "", "File to save the output to")
	flag.Parse()

	assemble(entryFile, outputFile)
}
