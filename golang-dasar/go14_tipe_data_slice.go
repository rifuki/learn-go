package main

import "fmt"

/** slice data type is most used in golang */
func main() {
	colors := [...]string{
		"scarlet",
		"indigo",
		"teal",
		"magenta",
		"turqoise",
		"amber",
	}
	fmt.Printf("list of colors: %v\n", colors)
	slice_from_colors := colors[4:6]

	fmt.Printf("first slice: %v\n", slice_from_colors[0])
	fmt.Printf("second slice: %s\n", slice_from_colors[1])

	slice_color_beginning_to_3 := colors[:3] /* <- slice from beginning to index 3 */
	fmt.Printf("slice_color_beggining_to_3: %v\n", slice_color_beginning_to_3)

	var slice_color_3_to_end []string = colors[3:] /* <- slice from index 3 to end array */
	fmt.Printf("slice_color_3_to_end: %v\n", slice_color_3_to_end)

	slice_all := colors[:] /* <- take all array value to slice */
	fmt.Printf("slice_all_colors: %v\n", slice_all)

	fmt.Println()

	/** function slice */
	days := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Printf("list of days: %v\n", days)
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Printf("list of days after reassignment: %v\n", days)
	fmt.Printf("list of daySlice1: %v\n", daySlice1)

	/** append slice */
	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Printf("list of daySlice2: %v\n", daySlice2)
	daySlice2[0] = "Ups"
	fmt.Printf("list of of daySlice2 reassignment: %v\n", daySlice2)
	fmt.Printf("list of days: %v\n", days)

	fmt.Println()

	/** create slice from scratch */
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Manggo"
	newSlice[1] = "Grape"
	// newSlice[2] = "Orange" /* <- this will error (use instead of append method for slice) */
	fmt.Printf("newSlice: %v\n", newSlice)
	fmt.Printf("length of newSlice: %d\n", len(newSlice))
	fmt.Printf("capacity of newSlice: %d\n", cap(newSlice))

	newSlice2 := append(newSlice, "Orange")
	fmt.Printf("newSlice2: %v\n", newSlice2)
	fmt.Printf("length of newSlice2: %d\n", len(newSlice2))
	fmt.Printf("capacity of newSlice2: %d\n", cap(newSlice2))

	/*** prove newSlice2 is refers newSlice */
	newSlice2[0] = "Apple"
	fmt.Printf("newSlice after reassignment newSlice2[0]: %v\n", newSlice)

	fmt.Println()

	/** copy slice */
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Printf("toSlice: %v\n", toSlice)

	fmt.Println()

	/** array vs slice */
	thisArray := [...]string{"one", "two", "three", "four"}
	thisSlice := []string{"action", "fantasy", "romance", "drama", "slice of life"}

	fmt.Printf("thisArray: %v\n", thisArray)
	fmt.Printf("thisSlice: %v\n", thisSlice)
}
