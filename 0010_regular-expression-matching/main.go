package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	// start string lengths
	sLen, pLen := len(s), len(p)

	// Initialize a 2D slice (dynamic programming table) with dimensions (sLen+1) x (pLen+1)
	// dp[i][j] will be true if s[:i] matches p[:j]
	dp := make([][]bool, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1)
	}

	// position 0,0 is always true
	// An empty pattern matches an empty string
	dp[0][0] = true

	// Fill in the dp table for cases where the pattern has '*' which can match an empty string
	for j := 2; j <= pLen; j++ {
		if p[j-1] == '*' {
			// If a '*' is found, check if dp[0][j-2] was true (pattern can ignore '*')
			dp[0][j] = dp[0][j-2]
		}
	}

	// Iterate over each character of the input string
	for i := 1; i <= sLen; i++ {

		// Iterate over each character of the pattern
		for j := 1; j <= pLen; j++ {

			// Check if the current characters of s and p match, or if the pattern has a '.'
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// If they match, carry forward the previous state
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			// If the pattern character is '*', we have two options:
			if p[j-1] == '*' {
				// 1. Treat '*' as matching zero occurrences of the previous character (dp[i][j-2])
				// 2. Treat '*' as matching one or more occurrences, so check the previous string character (dp[i-1][j])
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (p[j-2] == '.' || p[j-2] == s[i-1]))
				continue
			}

			// If characters don't match and there's no wildcard, mark as false
			dp[i][j] = false
		}
	}

	return dp[sLen][pLen]
}

// Failed Submission
//func isMatch(s string, p string) bool {
//	sIndex := 0
//	pIndex := 0
//
//	var preceding uint8
//	var sChar string
//	var pChar string
//	for sIndex < len(s) {
//		if pIndex == len(p) {
//			return false
//		}
//
//		sChar, pChar = string(s[sIndex]), string(p[pIndex])
//		_, _ = sChar, pChar
//
//		if p[pIndex] == '*' {
//			if preceding == '.' || s[sIndex] == preceding {
//				sIndex++
//				continue
//			}
//
//			pIndex++
//			continue
//		}
//
//		if s[sIndex] == p[pIndex] {
//			preceding = s[sIndex]
//			sIndex++
//			if sIndex < len(s) {
//				pIndex++
//			}
//
//			continue
//		}
//
//		if p[pIndex] == '.' {
//			preceding = '.'
//			sIndex++
//			pIndex++
//			continue
//		}
//
//		if len(p) > pIndex+1 && p[pIndex+1] == '*' {
//			pIndex += 2
//			continue
//		}
//
//		return false
//	}
//
//	if pIndex+1 < len(p) && p[pIndex] == '*' {
//		return preceding == p[pIndex+1]
//	}
//
//	// s=aaa, p=aaaa => false
//	if pIndex+1 < len(p) && p[pIndex+1] != '*' {
//		return false
//	}
//
//	return true
//}

func main() {
	// Test cases stored in a map
	testCases := map[[2]string]bool{
		{"aa", "a"}:                   false, // Expected: false
		{"aa", "a*"}:                  true,  // Expected: true
		{"ab", ".*"}:                  true,  // Expected: true
		{"aab", "c*a*b"}:              true,  // Additional case for variety
		{"mississippi", "mis*is*p*."}: false, // Additional case for complexity
		{"ab", ".*c"}:                 false, // Additional case for complexity
		{"aaa", "aaaa"}:               false, // Additional case for complexity
		{"aaa", "a*a"}:                true,  // Additional case for complexity
		{"aaa", "ab*a*c*a"}:           true,  // Additional case for complexity
	}

	// Iterating over test cases
	for input, expected := range testCases {
		s, p := input[0], input[1]
		result := isMatch(s, p)
		fmt.Printf("Input: s=%s, p=%s, Expected: %v, Output: %v\n", s, p, expected, result)
	}
}
