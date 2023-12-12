package main

import "fmt"

func main() {
	var person1 map[string]string = map[string]string{}
	person1["name"] = "Aozora"
	person1["age"] = "20"

	person2 := map[string]string{
		"name": "Rifuki",
		"age":  "22",
	}
	fmt.Printf("person2: %v\n", person2)
	fmt.Printf("person2 name: %s\n", person2["name"])
	fmt.Printf("person2 age: %s\n", person2["age"])
	fmt.Printf("person2 salah: %v\n", person2["salah"]) /* <- this will use default value if key map is missing */

	/** function map */
	books := make(map[string]string)
	books["title"] = "Go-Lang Book"
	books["author"] = "Robert Griesemer"
	books["wrong"] = "Ups"

	fmt.Printf("books: %v\n", books)
	delete(books, "wrong")
	fmt.Printf("books after deleted: %v\n", books)
}
