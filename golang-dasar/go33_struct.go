package main

import "fmt"

/** struct is a data template that is used to combine zero or more other data type in one unit */
/** data struct stored in field */
/** simply, struct is collection of fields */
/** cannot be used directly */

type Customer struct {
	Name, Address string /* <- naming convention is PascalCase */
	Age           int
}

func main() {
	var customer1 Customer
	fmt.Printf("customer1 before assignment: %v\n", customer1) /* <- automatically fill with default values */

	customer1.Name = "Rifuki"
	customer1.Address = "Solo"
	customer1.Age = 22
	fmt.Printf("customer1 after assignment: %v\n", customer1)

	fmt.Println()

	/** accessing one by one */
	fmt.Printf("customer1 name: %s\n", customer1.Name)
	fmt.Printf("customer1 age: %d\n", customer1.Age)
	fmt.Printf("customer1 address: %s\n", customer1.Address)

	fmt.Println()

	/** struct literals */
	customer2 := Customer{
		Name:    "John",
		Address: "No Law City",
		Age:     16,
	}
	fmt.Printf("customer2 name: %s\n", customer2.Name)
	fmt.Printf("customer2 age: %d\n", customer2.Age)
	fmt.Printf("customer2 address: %s\n", customer2.Address)

	fmt.Println()

	customer3 := Customer{"Hope", "Midgar", 17} /* <- must be sequential */
	fmt.Printf("customer3 name: %s\n", customer3.Name)
	fmt.Printf("customer3 age: %d\n", customer3.Age)
	fmt.Printf("customer3 address: %s\n", customer3.Address)
}
