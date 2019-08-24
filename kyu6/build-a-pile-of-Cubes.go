package kyu6

import (
	"math"
)

// https://www.codewars.com/kata/5592e3bd57b64d00f3000047/train/go

// FindNb takes a volume m
// returns the number of cubes to build pile
func FindNb(m int) int {

	volume := float64(m)
	root := math.Sqrt(volume)
	mod := int(volume) % int(root)

	// Check if Sqrt is multiple of volume
	if mod != 0 {
		return -1
	}

	n := 0
	i := 1
	// Loop till we find n
	for n <= int(root) {
		n += i

		if n == int(root) {
			return i
		}

		i++
	}

	// Default return value if no n is found
	return -1
}
