package main

import "fmt"

func main() {
	var nilaiAkhir = 80
	var nilaiAbsensi = 89

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusNilaiAbsensi bool = nilaiAbsensi >= 85

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsensi

	fmt.Println("Lulus :", lulus)
}
