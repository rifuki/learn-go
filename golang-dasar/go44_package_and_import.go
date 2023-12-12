package main

/** the place used to organize the program (like folder) */
/** package naming follows the name of the folder where the program code is stored */
/** use the `import` keyword to access program code that is in a different package */

import (
	"fmt"
	"golang-dasar/go44_helper"
)

func main() {
	result := go44_helper.SayHello("Rifuki")
	fmt.Printf("Result: %s\n", result)
}
