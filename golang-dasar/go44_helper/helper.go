package go44_helper

func SayHello(name string) string { /* <- must Capitalized function name to makes access modifier is public so outer package can access it.*/
	return "Hello " + name
}
