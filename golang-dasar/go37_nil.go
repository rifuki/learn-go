package main

import "fmt"

/** nil or null data types */
/** in golang when creating variable with certain data types, the default value will automatically be created */
/** in golang there is nil data, which is empty data */
/** nil can only used in several data types like `interface`, `function`, `map`, `slice`, `pointer`, and `channel` */

// func contoh(name string) string {
// 	if name == "" {
// 		return nil /* <- nil can't be used on string data type */
// 	} else {
// 		return name
// 	}
// }

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("masenko")
	if data == nil {
		fmt.Println(data)
	} else {
		fmt.Println(data)
	}
}
