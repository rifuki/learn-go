package main

import "fmt"

/** the ability to change the data type to the desired data type */
/** this feature is often used when we encounter empty interface data */
/** use .() notation */
/** type assertions returns 2 value (result conversion and result is conversion is succes) */

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	// resultString := result.(string) /* <- simpel ways if know data type before */
	// fmt.Printf("resultString: %s\n", resultString)

	// resultInt := result.(int) /* <- will return panic */
	// fmt.Printf("resultInt: %d\n", resultInt)

	/** safe ways using type assertions */
	switch value := result.(type) {
	case string:
		fmt.Printf("String %s\n", value)
	case int:
		fmt.Printf("Integer %d\n", value)
	default:
		fmt.Printf("Unknown %v\n", value)
	}
}
