package Select



func Selectu64(x uint64, i uint64) int {
	return TrailingZeros64(pdep(x, i))

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

func TrailingZeros64(x uint64) int {
	var result int = 0

	asm := `
    // Loop initialization
    MOVQ %[0], %[1] // result = x
    MOVQ $63, %[2] // pos = 63

    // Loop condition
.cond_check:
    CMPQ $0, %[1] // Check if result is zero
    JZ .done

    ANDQ $1, %[1] // Check the lowest bit
    JNZ .bit_not_set

    // Bit is not set
    JMP .bit_set

.bit_not_set:
    // Decrement pos
    DECQ %[2]

    // Shift right by one bit
    SHRQ $1, %[1]

    // Jump to the loop condition
    JMP .cond_check

.bit_set:
    // Store the position in result
    MOVQ %[2], %[1]
    RET

.done:
    `
	_, _, _, result = asm, x, 0, 0

	return result
}

