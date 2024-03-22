package main

import (
	"fmt"
	"go/praktikum/praktikumpackage"
	// "luaskeliling"
)

func main() {
	panjang := 5
	lebar := 3

	luas, keliling := praktikumpackage.HitungLuasKeliling(panjang, lebar)

	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)
}
