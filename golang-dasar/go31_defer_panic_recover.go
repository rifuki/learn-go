package main

/** defer, panic, and recover is like try-catch */

import "fmt"

/** `defer` is a keyword that is used to delay (defer) the execution of a statement until the function in which the statement resides has finished executing */
func logging() {
	fmt.Println("Finished call function")
}
func runApplication() {
	defer logging() /* <- function with defer keyword always executed even if there is an error code statement */
	fmt.Println("Run Application")
}

/** `panic` is a function to stop the program */
func endApp() {
	fmt.Println("end app")
	message := recover() /* <- recover keyword will catch the panic data/message and will not stop the program */
	fmt.Printf("an error occured: %s\n", message)
}
func runApp(error bool) {
	defer endApp()

	if error {
		panic("ups error") /* <- panic will force stop program */
	}
}

/** recover is used to catch panic data (combine in function that used defer keyword) */

func main() {
	runApplication()

	fmt.Println()

	runApp(true)

	fmt.Println()
}
