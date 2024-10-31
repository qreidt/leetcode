package main

import (
	"fmt"
	"slices"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var matrix [][]string

	// Montar matriz coluna por coluna
	//  x i0  i1  i2
	// j0 0	      4
	// j1 1   3   5
	// j2 2       6
	//
	//
	//  x i0  i1  i2  i3
	// j0 0	          6
	// j1 1       5   7
	// j2 2   4       8
	// j3 3           9
	//
	for i := 0; i < len(s); i++ {
		column := slices.Repeat([]string{""}, numRows)
		columnThreshold := numRows + (numRows - 2)
		columnOffset := i % columnThreshold // skip: 3

		// Build Full Column
		if columnOffset == 0 {

			// Make sure it doesn't overflow
			columnLength := numRows
			if len(s) < i+columnLength {
				columnLength = len(s) - i
			}

			// Add each char in the column
			for j := 0; j < columnLength; j++ {
				column[j] = string(s[i])

				if j < columnLength-1 {
					i = i + 1
				}
			}
		} else {
			// numRows = 3, offset = 3, threshold = 4 -> 1
			//
			// numRows = 4, offset = 4, threshold = 6 -> 2
			// numRows = 4, offset = 5, threshold = 6 -> 1
			columnLine := columnThreshold - columnOffset
			column[columnLine] = string(s[i])
		}

		matrix = append(matrix, column)
	}

	result := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(matrix); j++ {

			if matrix[j][i] != "" {
				result += matrix[j][i]
			}
		}
	}

	return result
}

//goland:noinspection ALL
func main() {
	//// Example 1
	s1 := "PAYPALISHIRING"
	numRows1 := 3
	fmt.Println("Output:", convert(s1, numRows1)) // Expected: "PAHNAPLSIIGYIR"

	// Example 2
	s2 := "PAYPALISHIRING"
	numRows2 := 4
	fmt.Println("Output:", convert(s2, numRows2)) // Expected: "PINALSIGYAHRPI"

	// Example 3
	s3 := "A"
	numRows3 := 1
	fmt.Println("Output:", convert(s3, numRows3)) // Expected: "A"
}
