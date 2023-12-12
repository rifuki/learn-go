package main

import "fmt"

/** like tuple in rust */

func getFullName() (string, string) {
	return "Mahoma", "Rifuki"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Printf("%s %s\n", firstName, lastName)

	/** ignore return value */
	_, lastNameIgnore := getFullName()
	fmt.Printf("%s\n", lastNameIgnore)
}
