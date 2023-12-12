package main

import "fmt"

func main() {
	/** Unary Operator */
	var j = 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)

	var name1 = "Chi"
	var name2 = "chi"

	result := name1 == name2

	fmt.Printf("'%s' == '%s': %t\n", name1, name2, result)
}
