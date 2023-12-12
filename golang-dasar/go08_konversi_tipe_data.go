package main

import "fmt"

func main() {
	nilai32 := 32768
	nilai64 := int64(nilai32)

	nilai16 := int16(nilai32) /* reached max int16 will result negative value. */

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	/** convert to string */
	name := "Rifuki"
	e_name := name[1]
	e_to_string := string(e_name)
	fmt.Println(e_to_string)
}
