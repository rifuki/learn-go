package main

import "fmt"

func main() {
	name := "rifuki"

	if name == "rifuki" {
		fmt.Println("hello rifuki")
	} else if name == "aceng" {
		fmt.Println("hello aceng")
	} else if name == "jamal" {
		fmt.Println("hello jamal")
	} else {
		fmt.Println("hello somebody")
	}

	/** if short statement */
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang.")
	} else {
		fmt.Println("nama sudah benar.")
	}
}
