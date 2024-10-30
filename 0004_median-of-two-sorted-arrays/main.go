package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	var n1 *int
	var n2 *int

	if len(nums1) > 0 {
		n1, nums1 = &nums1[0], nums1[1:]
	}

	if len(nums2) > 0 {
		n2, nums2 = &nums2[0], nums2[1:]
	}

	for {
		if n1 == nil && n2 == nil {
			break
		}

		if n1 != nil && n2 == nil {
			nums = append(nums, *n1)

			if len(nums1) == 0 {
				break
			}

			n1, nums1 = &nums1[0], nums1[1:]
			continue
		}

		if n2 != nil && n1 == nil {
			nums = append(nums, *n2)

			if len(nums2) == 0 {
				break
			}

			n2, nums2 = &nums2[0], nums2[1:]
			continue
		}

		if *n1 <= *n2 {
			nums = append(nums, *n1)

			if len(nums1) == 0 {
				n1 = nil
				continue
			}

			n1, nums1 = &nums1[0], nums1[1:]
		} else {
			nums = append(nums, *n2)

			if len(nums2) == 0 {
				n2 = nil
				continue
			}

			n2, nums2 = &nums2[0], nums2[1:]
		}
	}

	half := float64(len(nums))/2 - 1
	if len(nums)%2 == 0 {
		return float64(nums[int(half)]+nums[int(half+1)]) / 2
	}

	return float64(nums[int(math.Ceil(half))])
}

func main() {
	// Example 1
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("Output: %.5f\n", findMedianSortedArrays(nums1, nums2)) // Expected: 2.00000

	// Example 2
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Printf("Output: %.5f\n", findMedianSortedArrays(nums1, nums2)) // Expected: 2.50000

	// Example 3
	nums1 = []int{}
	nums2 = []int{1}
	fmt.Printf("Output: %.5f\n", findMedianSortedArrays(nums1, nums2)) // Expected: 2.50000
}
