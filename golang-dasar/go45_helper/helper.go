package go45_helper

var version = "1.0.0" /* <- can't accessed from outer package */
var Application = "golang"

/** can't accessed from outer package */
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
