package main

import "fmt"

/** very recommended use pointer in method */

type Man struct {
	Name      string
	isMarried bool
}

func (man *Man) Married() { /* <- its automatically use pointer without sending argument with reference operator (&)*/
	man.isMarried = !man.isMarried
}

func main() {
	man1 := Man{
		Name:      "Rifuki",
		isMarried: false,
	}
	fmt.Printf("is %s married: %v\n", man1.Name, man1.isMarried)
	man1.Married()
	fmt.Printf("is %s married: %v\n", man1.Name, man1.isMarried)
}
