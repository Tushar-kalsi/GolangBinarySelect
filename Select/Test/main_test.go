package Test

import (
	"Select/Select"
	"testing"
)

func TestSelectU64(t *testing.T) {
	testCases := []struct {
		x, i       uint64
		expected   int
	}{
		{0b0100110101101, 3, 0},  // Example test case 1
		{0b0100110101101, 5, 1},  // Example test case 2
		{0b0100110101101, 10, 0}, // Example test case 3
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		result := Select.Selectu64(tc.x, tc.i)
		if result != tc.expected {
			t.Errorf("select_u64(%d, %d) = %d; want %d", tc.x, tc.i, result, tc.expected)
		}
	}
}



func TestSelectu64(t *testing.T) {
    // Test with known inputs and expected results
    testCases := []struct {
        x, i      uint64
        expected  int
    }{
        {0b1010101010101010, 0b0000111100001111, 2}, // Example 1: x has two set bits, and i has four set bits
        {0b1111111111111111, 0b1111111111111111, 0}, // Example 2: x and i are both all set bits
        {0, 0b1000000000000000, 0}, // Example 3: x is 0, and i has only the most significant bit set
        {0, 0, 0}, // Example 4: Both x and i are 0
        {0b1100001100001100, 0b1100110000110001, 1}, // Example 5: x and i have different patterns of set bits
        {0b0101010101010101, 0b0101010101010101, 0}, // Example 6: x and i have alternating set and unset bits
        {0, 0xFFFFFFFFFFFFFFFF, 0}, // Example 7: x is 0, and i has all bits set
        {0xFFFFFFFFFFFFFFFF, 0, 63}, // Example 8: x has all bits set, and i is 0
        {0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0}, // Example 9: x and i are both all bits set
    }

    for _, tc := range testCases {
        result := Select.Selectu64(tc.x, tc.i)
        if result != tc.expected {
            t.Errorf("For x=%#x and i=%#x, expected %d, but got %d", tc.x, tc.i, tc.expected, result)
        }
    }
}
func TestTrailingZeros64(t *testing.T) {
	testCases := []struct {
		x        uint64
		expected int
	}{
		{0b0000000000001, 0}, // Example test case 1
		{0b0000000000010, 1}, // Example test case 2
		{0b0000000001000, 3}, // Example test case 3
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		result := Select.TrailingZeros64(tc.x)
		if result != tc.expected {
			t.Errorf("TrailingZeros64(%d) = %d; want %d", tc.x, result, tc.expected)
		}
	}
}
