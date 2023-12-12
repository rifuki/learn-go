package main

import "fmt"

/** interface is abstract data type */
/** an interface contains method definitions */
/** interface usually used as contracts (must be implemented) */

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(value HasName) {
	fmt.Printf("Hello %s\n", value.GetName())
}

type Animal struct {
	name string
}

func (animal Animal) GetName() string {
	return animal.name
}

func main() {
	person := Person{
		Name: "Rifuki",
	}

	SayHello(person)

	fmt.Println()

	animal := Animal{
		name: "Hikari",
	}
	SayHello(animal)
}
