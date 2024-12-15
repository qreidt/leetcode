package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Place numbers in their correct positions
	for i := 0; i < n; i++ {
		for 1 <= nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// Identify first missing positive
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {

	testCases := []struct {
		values []int
		result uint
	}{
		//{values: []int{1, 2, 0}, result: 3},
		{values: []int{3, 4, -1, 1}, result: 2},
		{values: []int{7, 8, 9, 11, 12}, result: 1},
	}

	for _, tc := range testCases {
		fmt.Println("---------------------------------")
		result := firstMissingPositive(tc.values)
		fmt.Printf("INPUT: %v\n", tc.values)
		fmt.Printf("EXPECTED: %d\n", tc.result)
		fmt.Printf("RESULT: %d\n\n", result)
	}

}
