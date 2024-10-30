package main

import (
	"fmt"
	"slices"
)

func lengthOfLongestSubstring(s string) int {
	var bestSequenceLength int
	var currentSequence []uint8
	sequenceStart := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if !slices.Contains(currentSequence, char) {
			currentSequence = append(currentSequence, char)
			continue
		}

		if len(currentSequence) > bestSequenceLength {
			bestSequenceLength = len(currentSequence)
		}

		currentSequence = []uint8{}
		i = sequenceStart
		sequenceStart += 1
	}

	if len(currentSequence) > bestSequenceLength {
		return len(currentSequence)
	}

	return bestSequenceLength
}

//goland:noinspection SpellCheckingInspection
func main() {
	// Example 1
	//s1 := "abcabcbb"
	//fmt.Println("Output:", lengthOfLongestSubstring(s1)) // Expected: 3
	//
	//// Example 2
	//s2 := "bbbbb"
	//fmt.Println("Output:", lengthOfLongestSubstring(s2)) // Expected: 1
	//
	//// Example 3
	//s3 := "pwwkew"
	//fmt.Println("Output:", lengthOfLongestSubstring(s3)) // Expected: 3

	// Example 4
	s4 := "dvdf"
	fmt.Println("Output:", lengthOfLongestSubstring(s4)) // Expected: 3
}
