package main

import "fmt"

/** (*) like dereferencing in rust */

type Address struct {
	city    string
	state   string
	country string
}

func main() {
	address1 := Address{
		city:    "Surakarta",
		state:   "Central Java",
		country: "Indonesia",
	}
	memory_address1 := &address1

	var address2 *Address = &address1
	address2.city = "Pekalongan"

	// address2 = &Address{"Bandung", "West Java", "Indonesia"} /* <- will make new object */
	*address2 = Address{"Bandung", "West Java", "Indonesia"} /* <- change the referenced data*/

	fmt.Printf("address1: %v | memory: %v\n", address1, &memory_address1)
	fmt.Printf("address2: %v | memory: %v\n", address2, &address2)
}
