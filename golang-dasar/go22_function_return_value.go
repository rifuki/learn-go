package main

import "fmt"

func getHello(name string) string {
	return "hello" + name
}

func main() {
	result := getHello("rifuki")
	fmt.Println(result)

	fmt.Println(getHello("john"))
}
