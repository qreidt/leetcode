package main

import "fmt"

func main() {
	// Example 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println("Output:", twoSum(nums1, target1)) // Expected: [0, 1]

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println("Output:", twoSum(nums2, target2)) // Expected: [1, 2]

	// Example 3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println("Output:", twoSum(nums3, target3)) // Expected: [0, 1]
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
