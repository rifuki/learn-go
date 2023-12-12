package main

import "fmt"

/** golang as default passing as value */
/** send variable to function, method, or other variable sending a duplicate value (like clone, copy in rust) */

type Address struct {
	city    string
	state   string
	country string
}

/** a pointer is the ability to create a reference to a data location in the same memory, without duplicating existing data. */
/** with pointer can create a pass by reference */

func main() {
	address1 := Address{
		city:    "Surakarta",
		state:   "Central Java",
		country: "Indonesia",
	}
	memory_address1 := &address1

	/** pass by value */
	address2 := address1 /* <- clone from address1 */
	address2.city = "Pekalongan"
	memory_address2 := &address2

	/** pass by reference (& notation) */
	var address3 *Address = &address1 /* <- reference to address1 */
	address3.city = "Sragen"
	address4 := &address1 /* <- reference to address1 */
	address4.city = "Boyolali"

	fmt.Printf("address1: %v | memory: %v\n", address1, &memory_address1)
	fmt.Printf("address2: %v | memory: %v\n", address2, &memory_address2)
	fmt.Printf("address3: %v | memory: %v\n", address3, &address3)
	fmt.Printf("address4: %v | memory: %v\n", address4, &address4)
}
