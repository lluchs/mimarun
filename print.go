// Printing utilities

package main

import (
	"fmt"

	"github.com/lluchs/mima"
)

// Prints a program's pointers and instructions.
func PrintProgram(program *mima.Program) {
	fmt.Println("Pointer")
	for mark, pointer := range program.Marks {
		fmt.Printf("%-10s 0x%06X = %d\n", mark, pointer, pointer)
	}

	fmt.Println("\nInstructions")
	for address, instruction := range program.Instructions {
		fmt.Printf("0x%06X = %s(%s)\n", address, instruction.Op, instruction.Argument)
	}
}

// Print all memory locations which are not 0.
func PrintMem(mem []uint32) {
	for pos, content := range mem {
		if content != 0 {
			fmt.Printf("0x%06X: 0x%06X\n", pos, content)
		}
	}
}
