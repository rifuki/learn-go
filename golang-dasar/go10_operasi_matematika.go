package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	c := a + b
	fmt.Println(c)

	/** Augmented Assignment */
	i := 10
	fmt.Println(i)
	i += 1
	fmt.Println(i)
	i += 2
	fmt.Println(i)
}
