package main

import "fmt"

/** empty interface allows provide any data types */
/** interface {} */
/** any is type declaration for empty interface */

func Ups() any {
	// return 1
	// return "satu" /* <- can return all data types */
	return true
}

func main() {
	var kosong = Ups()
	fmt.Printf("ups: %v\n", kosong)
}
