package main

import (
	"C"
	"fmt"
	"math/bits"
)

// func pdep_u64(x, i uint64) uint64 {
//     // Simulate the "pdep" operation using bitwise shifting and masking
//     mask := uint64(1) << i
//     result := x & mask

//     return result
// }

func select_u64(x uint64, i uint64) int {

	return bits.TrailingZeros64(pdep(x, i))

}

func main() {
	x := uint64(0b0100110101101) // Example binary sequence
	i := uint64(3)               // Select the 3rd bit

	selectedBit := select_u64(x, i)
	fmt.Printf("Selected bit at position %d: %b\n", i, selectedBit)
}

func pdep(x, mask uint64) uint64 {
	result := uint64(0)

	asm := `
        MOV x0, %[x]
        MOV x1, %[mask]
        MRS x2, SPSR_EL1

        // Use ARM64's "EOR" instruction to simulate the "pdep" operation
        CLZ x3, x1
        EOR x4, x0, x0, LSR x3
        EOR x5, x0, x0, LSR x3, LSL x3
        EOR x6, x0, x0, LSR x3, LSL x3, LSL x3
        ORR %[result], x4, x5
        ORR %[result], %[result], x6

        MSR SPSR_EL1, x2
    `
	asmWrapper(asm, x, mask, &result)

	return result
}

func asmWrapper(asm string, i, x uint64, result *uint64)
