package main

import "fmt"

func main() {
	name := "Darmono"

	if name == "Dina" {
		fmt.Println("Hello Dina")
	} else if name == "Euis" {
		fmt.Println("Hello Euis")
	} else {
		fmt.Println("Halo dek, boleh kenalan ndak?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
