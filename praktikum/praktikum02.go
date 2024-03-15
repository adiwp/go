package main

import "fmt"

func main() {
	var a int = 10
	var b int = 5

	var hasilPenjumlahan = a + b
	var hasilPengurangan = a - b
	var hasilPerkalian = a * b
	var hasilPembagian = a / b
	var sisaBagi = a % b

	fmt.Println("Hasil penjumlahan:", hasilPenjumlahan)
	fmt.Println("Hasil pengurangan:", hasilPengurangan)
	fmt.Println("Hasil perkalian:", hasilPerkalian)
	fmt.Println("Hasil pembagian:", hasilPembagian)
	fmt.Println("Sisa bagi:", sisaBagi)
}
