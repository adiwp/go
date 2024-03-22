package main

import (
	"fmt"
	"math"
	"progo/ratarata" // Mengimpor package
)

func main() {
	// Data nilai siswa
	nilaiSiswa := []int{80, 75, 90, 85, 60}

	// Menghitung rata-rata dengan fungsi dari package
	rataRata := ratarata.HitungRataRata(nilaiSiswa)

	// Menampilkan hasil
	fmt.Println("Rata-rata nilai siswa:", rataRata)

	/*******/
	// anonymous function
	/*******/
	jariJari := 5.0

	// Menghitung luas dan keliling dengan anonymous function
	luas := func(r float64) float64 {
		return math.Pi * r * r
	}(jariJari)

	keliling := func(r float64) float64 {
		return 2 * math.Pi * r
	}(jariJari)

	// Menampilkan hasil
	fmt.Println("Luas lingkaran:", luas)
	fmt.Println("Keliling lingkaran:", keliling)

	// closure
	// Batas nilai untuk setiap huruf
	batasNilai := map[string]int{
		"A": 90,
		"B": 80,
		"C": 70,
		"D": 60,
		"E": 0,
	}

	// Closure untuk menentukan nilai huruf
	getNilaiHuruf := func(nilaiUjian int) string {
		for huruf, batas := range batasNilai {
			if nilaiUjian >= batas {
				return huruf
			}
		}
		return "E"
	}

	// Menentukan nilai huruf untuk beberapa nilai ujian
	nilaiUjian := []int{85, 75, 95, 55, 65}
	for _, nilai := range nilaiUjian {
		fmt.Println("Nilai ujian", nilai, ":", getNilaiHuruf(nilai))
	}
}
