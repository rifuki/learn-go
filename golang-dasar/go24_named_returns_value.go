package main

import "fmt"

/** declaring in returns parameter function */
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Osma"
	lastName = "Sora"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
