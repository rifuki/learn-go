package main

import "fmt"

/** function is first class citizen has privileges */

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye /* <- function as value */
	fmt.Println(goodbye("rifuki"))
}
