package main

import (
	"fmt"
	"os"

	"github.com/lluchs/mima"
)

func main() {
	argc := len(os.Args)
	if argc != 2 && argc != 3 {
		fmt.Println("Usage: mimarun <filename>")
		return
	}
	var filename, command string
	if argc == 2 {
		command = "run"
		filename = os.Args[1]
	} else {
		command = os.Args[1]
		filename = os.Args[2]
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

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
	var mem []uint32
	switch command {
	case "analyze":
		mem, err = bytecode.Debug(analyze(bytecode))
	default:
		mem, err = bytecode.Run()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the resulting memory.
	PrintMem(mem)

}
