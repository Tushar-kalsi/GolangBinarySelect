package Select

import (
	"Select/Select"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSelectu64AverageTime(b *testing.B) {
	// Seed the random number generator for generating random input values
	rand.Seed(time.Now().UnixNano())

	var totalDuration time.Duration

	for i := 0; i < b.N; i++ {
		// Generate random uint64 values for x and i
		x := rand.Uint64()
		mask := rand.Uint64()

		// Record the start time
		startTime := time.Now()

		// Call the function to be benchmarked
		_ = Select.Selectu64(x, mask)

		// Calculate the elapsed time and add it to the total duration
		elapsedTime := time.Since(startTime)
		totalDuration += elapsedTime
	}

	// Calculate the average time
	averageTime := totalDuration / time.Duration(b.N)

	// Print the average time (in nanoseconds) to the console
	b.Logf("Average time: %v", averageTime)
}
