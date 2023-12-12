package main

import "fmt"

/** method is a function that attached to a struct */

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Printf("Hello %s, My Name is %s\n", name, customer.Name)
}

func main() {
	var customer Customer = Customer{
		Name:    "Rifuki",
		Address: "Solo",
		Age:     22,
	}
	customer.sayHello("Foo")
}
