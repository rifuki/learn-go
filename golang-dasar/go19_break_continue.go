package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("break looping-%d\n", i)
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("continue looping-%d\n", i)
	}
}
