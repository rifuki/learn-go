package main

import "fmt"

/** to create a variable, start with the `var` keyword, followed by the variable name1 and data type. */
func main() {
	var name1 string

	name1 = "rifuki"
	fmt.Println(name1)

	name1 = "wikwok"
	fmt.Println(name1)
	fmt.Println()

	/*** data types are automatically inferred */
	var name2 = "john"
	fmt.Println(name2)

	name2 = "smith"
	fmt.Println(name2)
	fmt.Println()

	/*** create variable with `:=` (best practice) */
	name3 := "suzuki"
	fmt.Println(name3)

	name3 = "hope"
	fmt.Println(name3)
	fmt.Println()

	/*** deklarasi multiple variable */
	var (
		firstName = "Mahoma"
		lastName  = "Rifuki"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
