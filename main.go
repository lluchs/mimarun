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
		fmt.Printf("%-10s 0x%06X = %d\n", mark, pointer, pointer)
	}

	fmt.Println("\nInstructions")
	for address, instruction := range program.Instructions {
		fmt.Printf("0x%06X = %s(%s)\n", address, instruction.Op, instruction.Argument)
	}

	fmt.Println("\n\nAssembling...\n")
	bytecode, err := program.Assemble()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the assembled code.
	fmt.Printf("Start: 0x%06X\n", bytecode.Start)
	printMem(bytecode.Mem)

	fmt.Println("\n\nRunning...\n")
	mem, err := bytecode.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the resulting memory.
	printMem(mem)

}

// Print all memory locations which are not 0.
func printMem(mem []uint32) {
	for pos, content := range mem {
		if content != 0 {
			fmt.Printf("0x%06X: 0x%06X\n", pos, content)
		}
	}
}
