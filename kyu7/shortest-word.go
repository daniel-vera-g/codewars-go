package kyu7

import (
	"strings"
)

// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9
// FindShort returns the shortest number in for a given array of words
func FindShort(s string) int {

	// convert string into array
	sArray := strings.Fields(s)

	var shortestWordLength int = len(sArray[0])

	// find the shortest word
	for _, word := range sArray {
		if len(word) < shortestWordLength {
			shortestWordLength = len(word)
		}
	}

	return shortestWordLength
}
