package main

import "fmt"

/** factorial with for loop */
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {
	fmt.Printf("factorial for_loop: %d\n", factorialLoop(3))

	fmt.Printf("factorial recursive function: %d\n", factorialRecursive(3))
}
