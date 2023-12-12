package main

import "fmt"

func main() {
	const firstName string = "Mahoma"
	const lastName = "Rifuki"

	// firstName = "Aceng" /* <- it will error */
	// lastName = "Firmawan" /* <- it will error */

	/** Deklarasi Multiple Constant */
	const (
		phi1 = 3.14
		phi2 = 22 / 7
	)

	fmt.Printf("Nama: %s, Umur: %d\n", lastName, phi2)
}
