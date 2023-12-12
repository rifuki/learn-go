package main

import "fmt"

func main() {
	counter := 0
	fmt.Printf("Counter: %d\n", counter)
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Printf("Counter: %d\n", counter)
}
