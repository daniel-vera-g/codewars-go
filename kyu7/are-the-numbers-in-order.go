package kyu7

// InAscOrder gets an array of integers as input
// return a bool if the array is in ascending order
func InAscOrder(numbers []int) bool {

	var status bool
	var point int

	if numbers[0] == 1 && len(numbers) == 1 {
		return true
	}

	// Loop through array and check ascending rule
	for i, number := range numbers {

		if (len(numbers) - 1) != i {
			point = i + 1
		}

		if number <= numbers[point] {
			status = true
		} else {
			return false
		}

	}

	return status
}
