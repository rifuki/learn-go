package main

import (
	"fmt"
	"golang-dasar/go45_helper"
)

func main() {
	result := go45_helper.SayHello("Smith")
	fmt.Println(result)

	// fmt.Printf("version: %d\n", go45_helper.version) /* <- can't access it */
	// fmt.Printf("sayGoodBye: %s\n", go45_helper.sayGoodBye("rifuki")) /* <- can't access it */
	fmt.Printf("application: %s\n", go45_helper.Application)
}
