package main

import (
	"fmt"

	"github.com/lluchs/mima"
)

// Returns a MIMA debug function for analyzing program execution.
func analyze(bytecode *mima.Bytecode) mima.DebugFunc {
	lastMem := bytecode.Mem
	lastIAR := bytecode.Start
	i := 1
	fmt.Println("Step |   IAR    |    IR    |   Akku   | Memory Changes")
	fmt.Println("------------------------------------------------------")
	return func(state *mima.State) {
		// Compare memory in order to find changes.
		changed := ""
		for pos, val := range state.Mem {
			if val != lastMem[pos] {
				changed = fmt.Sprintf("0x%06X = %d = 0x%X", pos, val, val)
				// There can only be one change per cycle.
				break
			}
		}
		// Output
		fmt.Printf("%4d | 0x%06X | 0x%06X | 0x%6X | %s\n", i, lastIAR, state.IR, state.Akku, changed)
		// Copy the memory to be able to compare again.
		lastMem = make([]uint32, len(lastMem))
		copy(lastMem, state.Mem)
		lastIAR = state.IAR
		i++
	}
}
