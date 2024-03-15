package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "Adi",
		"alamat": "Pondok Indah",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["alamat"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Bambang Riono"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)

}
