package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Printf("looping-%d\n", counter)
		counter++
	}

	fmt.Println("finish")
	fmt.Println()

	/** for with statement */
	for init_counter := 1; init_counter <= 10; init_counter++ {
		fmt.Printf("looping with statement-%d\n", init_counter)
	}
	fmt.Println("finish looping with statement.")
	fmt.Println()

	/** for range (iterate over all data collection like Array, Slice, Map) */
	pokemons := []string{"Pikachu", "Eevee", "Mew", "Celebi"}
	// for i := 0; i < len(pokemons); i++ { /* <- usually ways (not recommended) */
	for index, name := range pokemons { /* <- use for range instead for statement (best practice) */
		fmt.Printf("pokemon-%d name: %s\n", index+1, name)
	}
}
