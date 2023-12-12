package main

/** error make normal error */

import (
	"errors"
	"fmt"
)

func Divide(nilai float32, pembagi float32) (float32, error) {
	if pembagi == 0 {
		return 0, errors.New("division by zero.")
	}
	return nilai / pembagi, nil
}

func main() {
	result1, error1 := Divide(3, 0)
	if error1 != nil {
		fmt.Printf("Error1: %s\n", error1.Error())
	} else {
		fmt.Printf("Result1: %d\n", result1)
	}

	fmt.Println()

	result2, error2 := Divide(1000, 3)
	if error2 != nil {
		fmt.Printf("Error2: %s\n", error2.Error())
	} else {
		fmt.Printf("Result2: %f\n", result2)
	}
}
