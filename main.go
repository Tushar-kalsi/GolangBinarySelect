package main

import (
	"C"
	"fmt"
)
import "Select/Select"


func main() {
	x := uint64(0b0100110101101) // Example binary sequence
	i := uint64(3)               // Select the 3rd bit

	selectedBit := Select.Selectu64(x, i)
	fmt.Printf("Selected bit at position %d: %b\n", i, selectedBit)
}
