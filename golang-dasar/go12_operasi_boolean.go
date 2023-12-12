package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi >= 80

	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Printf("Lulus: %t\n", lulus)
}
