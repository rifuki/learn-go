package main

import (
	"fmt"
	"golang-dasar/go46_database"
	_ "golang-dasar/go46_internal" /* <- use blank identifier (_) to only run the function without use it in main functio */
)

func main() {
	fmt.Printf("database: %s\n", go46_database.GetDatabase())
}
