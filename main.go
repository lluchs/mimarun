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
	PrintProgram(program)

	fmt.Println("\n\nAssembling...\n")
	bytecode, err := program.Assemble()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the assembled code.
	fmt.Printf("Start: 0x%06X\n", bytecode.Start)
	PrintMem(bytecode.Mem)

	fmt.Println("\n\nRunning...\n")
	mem, err := bytecode.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the resulting memory.
	PrintMem(mem)

}
