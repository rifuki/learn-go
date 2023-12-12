package main

import "fmt"

/** as default parameter function is pass by value */

type Address struct {
	city    string
	state   string
	country string
}

func changeCountryToIndonesia(address *Address) {
	address.country = "Indonesia"
}

func main() {
	address := Address{}
	changeCountryToIndonesia(&address)

	fmt.Printf("address country: %s\n", address.country)
}
