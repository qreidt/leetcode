package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// Negative numbers and numbers ending in 0 (except 0 itself) are not palindromes.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	return x == reversedHalf || x == reversedHalf/10
}

// Submit 1
//func isPalindrome(x int) bool {
//	str := strconv.Itoa(x)
//
//	half := len(str) / 2
//	_ = half
//
//	j := len(str) - 1
//	for i := 0; i <= half; i++ {
//		if str[i] != str[j] {
//			return false
//		}
//
//		j--
//	}
//
//	return true
//}

func main() {

	// Test cases stored in a map
	testCases := map[int]bool{
		121:  true,  // Expected: true
		-121: false, // Expected: false
		10:   false, // Expected: false
	}

	// Iterating over test cases
	for input, expected := range testCases {
		result := isPalindrome(input)
		fmt.Printf("Input: %d, Expected: %v, Output: %v\n", input, expected, result)
	}
}
