package main

import (
	"fmt"
	"os"

	"github.com/lluchs/mima"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: mimarun <filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Parsing...\n")
	program, err := mima.Parse(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the program.
	fmt.Println("Pointer")
	for mark, pointer := range program.Marks {
		fmt.Println(mark, pointer)
	}

	fmt.Println("Instructions")
	for address, instruction := range program.Instructions {
		fmt.Printf("%s = %s(%s)\n", address, instruction.Op, instruction.Argument)
	}
}
