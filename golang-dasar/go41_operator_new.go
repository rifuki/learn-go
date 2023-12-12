package main

import "fmt"

/** the new keyword can also be used to create pointers */
/** new keyword return pointer to empty data, not initial data */

type Address struct {
	city    string
	state   string
	country string
}

func main() {
	var alamat1 *Address = &Address{}
	var alamat2 *Address = alamat1
	alamat2.country = "Singapore"

	alamat3 := new(Address)
	alamat4 := alamat3

	alamat4.country = "Indonesia"

	fmt.Printf("alamat1: %v\n", alamat1)
	fmt.Printf("alamat2: %v\n", alamat2)
	fmt.Printf("alamat3: %v\n", alamat3)
	fmt.Printf("alamat4: %v\n", alamat4)
}
