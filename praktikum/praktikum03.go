package main

import "fmt"

func main() {
	var a int = 10
	var b int = 5

	var isSama = a == b
	var isTidakSama = a != b
	var isKurangDari = a < b
	var isKurangDariSamaDengan = a <= b
	var isLebihDari = a > b
	var isLebihDariSamaDengan = a >= b

	fmt.Println("Apakah sama?:", isSama)
	fmt.Println("Apakah tidak sama?:", isTidakSama)
	fmt.Println("Apakah a kurang dari b?:", isKurangDari)
	fmt.Println("Apakah a kurang dari atau sama dengan b?:", isKurangDariSamaDengan)
	fmt.Println("Apakah a lebih dari b?:", isLebihDari)
	fmt.Println("Apakah a lebih dari atau sama dengan b?:", isLebihDariSamaDengan)

}
