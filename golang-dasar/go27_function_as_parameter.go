package main

import "fmt"

type Filter func(string) string /* <- type declarations (like alias) */

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Printf("Hello %s\n", filter(name))
}

func spamFilter(name string) string {
	if name == "bitch" {
		return "....."
	}
	return name
}

func main() {
	sayHelloWithFilter("rifuki", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("bitch", filter)
}
