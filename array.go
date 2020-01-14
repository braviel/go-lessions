package main

// Sum returns sum of integer array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll return arrays contains sum of input arrays
func SumAll(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//SumAllTails return SumAll of tails of given slices
func SumAllTails(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
