package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

//go:noinline
func select_(x uint64, i uint64) int {
	return bits.TrailingZeros64(pdepu64(1<<i, x))
}

func pdepu64(i uint64, x uint64) uint64 {
	return (x >> i) & 1
}

func main2() {

	// Binary representation of the number
	binaryStr := "0100110101101"

	// Parse the binary string to a uint64
	val, err := strconv.ParseUint(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	ans := select_(val, 3)

	fmt.Println("Answer is ", ans)

	// Print the uint64 value
	fmt.Printf("Binary: %s\nUint64: %d\n", binaryStr, val)

}
