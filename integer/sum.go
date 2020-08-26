package integer

// Sum receives two integers and returns the sum of them
func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll is
func SumAll(numbers ...[]int) []int {
	var sum []int
	for _, number := range numbers {
		sum = append(sum, Sum(number))
	}

	return sum
}
