package GoSelect

import (
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
		result := select_u64(tc.x, tc.i)
		if result != tc.expected {
			t.Errorf("select_u64(%d, %d) = %d; want %d", tc.x, tc.i, result, tc.expected)
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
		result := TrailingZeros64(tc.x)
		if result != tc.expected {
			t.Errorf("TrailingZeros64(%d) = %d; want %d", tc.x, result, tc.expected)
		}
	}
}
