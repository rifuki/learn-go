package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "black"
	colors[1] = "grey"
	colors[2] = "white"

	fmt.Println(colors[0])
	fmt.Println(colors[1])
	fmt.Println(colors[2])

	/** declare direct array */
	var rgb = [3]int{
		250,
		120,
		/* <- default value */
	}
	fmt.Println(rgb)

	/** function array */
	fmt.Printf("length of rgb array: %d\n", len(rgb))
	rgb[2] = 999999999
	fmt.Printf("third element of rbg array: :%d\n", rgb[2])

	/** array without fixed length (must be declared immediately) */
	var my_array = [...]int{
		1,
		3,
		5,
		7,
		9,
		11,
		13,
		15,
		17,
	}
	fmt.Printf("my_array: %v\n", my_array)
}
