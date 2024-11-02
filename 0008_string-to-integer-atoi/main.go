package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	values := make([]uint8, len(s))

	valuesIndex := 0
	positiveOrNegative := 1
	isReadingValues := false
	hasLeadingZeros := false

	for _, charInt := range s {
		isNumericChar := charInt >= 48 && charInt < 58

		// Stop reading
		if ((isReadingValues || hasLeadingZeros) && !isNumericChar) || charInt >= 58 {
			break
		}

		// stop on "."
		if charInt == 46 {
			break
		}

		// Set negative and start reading
		if !isReadingValues && charInt == 45 {
			isReadingValues = true
			positiveOrNegative = -1
			continue
		}

		// Is positive
		if !isReadingValues && charInt == 43 {
			isReadingValues = true
			continue
		}

		// ignore leading zeros
		if valuesIndex == 0 && charInt == 48 {
			hasLeadingZeros = true
			continue
		}

		// skip if not numeric
		if !isReadingValues && !isNumericChar {
			if hasLeadingZeros {
				break
			}
			continue
		}

		values[valuesIndex] = uint8(charInt - 48)
		valuesIndex++
		isReadingValues = true
	}

	var result int = int(values[0]) * positiveOrNegative
	if len(values) > 1 {
		for j := 1; j < valuesIndex; j++ {
			result = result*10 + int(values[j])*positiveOrNegative

			// Validate if result is still below int32 threshold
			if result > math.MaxInt32 {
				return math.MaxInt32
			}

			if result < math.MinInt32 {
				return math.MinInt32
			}
		}
	}

	return result
}

func main() {
	// Example 1
	s1 := "42"
	fmt.Println("Output:", myAtoi(s1)) // Expected: 42

	//Example 2
	s2 := "   -042"
	fmt.Println("Output:", myAtoi(s2)) // Expected: -42

	// Example 3
	s3 := "1337c0d3"
	fmt.Println("Output:", myAtoi(s3)) // Expected: 1337

	// Example 4
	s4 := "0-1"
	fmt.Println("Output:", myAtoi(s4)) // Expected: 0

	// Example 5
	s5 := "words and 987"
	fmt.Println("Output:", myAtoi(s5)) // Expected: 0

	// Example 6
	s6 := "-91283472332"
	fmt.Println("Output:", myAtoi(s6)) // Expected: -2147483648

	// Example 7
	s7 := ".1"
	fmt.Println("Output:", myAtoi(s7)) // Expected: 0

	// Example 8
	s8 := "9223372036854775808"
	fmt.Println("Output:", myAtoi(s8)) // Expected: 2147483647

	// Example 9
	s9 := "2147483648"
	fmt.Println("Output:", myAtoi(s9)) // Expected: 2147483647

}
