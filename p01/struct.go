package main

import "fmt"

type Customer struct {
	Nama, Alamat string
	Umur         int
}

func (customer Customer) sayHello(nama string) {
	fmt.Println("Halo", nama, "my name is", customer.Nama)
}

func main() {
	var adi Customer
	fmt.Println(adi)

	adi.Nama = "Adi Wahyu"
	adi.Alamat = "Indonesia"
	adi.Umur = 17
	fmt.Println(adi)
	fmt.Println(adi.Nama)
	fmt.Println(adi.Alamat)
	fmt.Println(adi.Umur)

	joko := Customer{
		Nama:   "Joko",
		Alamat: "Indonesia",
		Umur:   30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	adi.sayHello("Agus")
	joko.sayHello("Agus")
}
