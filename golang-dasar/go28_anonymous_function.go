package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Printf("you are blocked. %s!\n", name)
	} else {
		fmt.Printf("welcome, %s\n", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "bitch"
	}

	registerUser("rifuki", blacklist)

	registerUser("bitch", func(name string) bool {
		return name == "bitch"
	})
}
