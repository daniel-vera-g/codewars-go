package kyu8

// https://www.codewars.com/kata/sum-of-positive/train/go

// PositiveSum returns the summ of all positive numbers in an array
func PositiveSum(numbers []int) int {

	// Sum of positive numbers
	var numbersSum int

	for _, e := range numbers {
		// Check for positive numbers
		if e > 0 {
			numbersSum += e
		}
	}

	return numbersSum
}
