package main

import "fmt"

func main() {
	dataMahasiswa := map[string]string{
		"123456": "Budi - Teknik Informatika",
		"789012": "Ani - Akuntansi",
		"345678": "Cici - Manajemen",
	}

	// Menampilkan daftar nama
	fmt.Println("Daftar nama mahasiswa:")
	for nim, data := range dataMahasiswa {
		fmt.Println("-", nim, "-", data)
	}

	// Mencari data berdasarkan NIM
	var nim string
	fmt.Print("Masukkan NIM: ")
	fmt.Scanf("%s", &nim)

	data, ok := dataMahasiswa[nim]
	if ok {
		fmt.Println("Data ditemukan:", data)
	} else {
		fmt.Println("Data mahasiswa dengan NIM tersebut tidak ditemukan")
	}
}
