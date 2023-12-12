package main

import "fmt"

func main() {
	fruit := "durian"

	switch fruit {
	case "apple":
		fmt.Printf("fruit: %s\n", fruit)
	case "durian":
		fmt.Printf("fruit: %s\n", fruit)
	default:
		fmt.Println("another fruit")
	}

	/** short statement */
	switch length := len(fruit); length > 5 {
	case true:
		fmt.Println("nama fruit terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	/** switch without condition */
	length := len(fruit)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama fruit sudash benar")
	}
}
