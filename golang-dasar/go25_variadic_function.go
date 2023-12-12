package main

import "fmt"

/** like rest parameter in javascript (last parameter can accept more than one input (like array)) */

func sumAllVariadic(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	totalVariadic := sumAllVariadic(10, 10, 10, 10, 10)
	fmt.Printf("total (variadic): %d\n", totalVariadic)

	totalSlice := sumAllSlice([]int{1, 1, 1, 1, 1})
	fmt.Printf("total (slice): %d\n", totalSlice)

	slice_numbers := []int{2, 2, 2, 2, 2}
	convertedSliceToVariadic := sumAllVariadic(slice_numbers...)
	fmt.Printf("total convertedSliceToVariadic: %d\n", convertedSliceToVariadic)
}

func sumAllSlice(numbers []int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}
