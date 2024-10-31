package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		// Validate if result is still below int32 threshold
		if (result > math.MaxInt32/10) || (result < math.MinInt32/10) {
			return 0
		}

		// pop last digit
		pop := x % 10
		x /= 10

		// Append digit to the reversed number
		result = result*10 + pop
	}

	return result
}

func main() {
	// Example 1
	x1 := 123
	fmt.Println("Output:", reverse(x1)) // Expected: 321

	//Example 2
	x2 := -123
	fmt.Println("Output:", reverse(x2)) // Expected: -321

	// Example 3
	x3 := 120
	fmt.Println("Output:", reverse(x3)) // Expected: 21

	// Example 4
	x4 := 1534236469
	fmt.Println("Output:", reverse(x4)) // Expected: 0
}
