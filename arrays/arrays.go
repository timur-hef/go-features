package arrays

func Sum(numbers []int) int {
	var total int

	for _, number := range numbers {
		total += number
	}

	return total
}

func SumAll(numArrays ...[]int) []int {
	results := make([]int, 0, len(numArrays)) // declare capacity, since we it will depend on the numArrays length

	for _, array := range numArrays {
		var total int

		for _, number := range array {
			total += number
		}

		results = append(results, total)
	}

	return results
}
