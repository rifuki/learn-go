package main

import "fmt"

type NoKTP string

func main() {

	var ktpAceng NoKTP = "3937501020100013"
	ktpSupri := NoKTP("2893423108322239")
	fmt.Println(ktpAceng)
	fmt.Println(ktpSupri)
	fmt.Println(NoKTP("1294984379820928"))
}
