package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// for range loop
	var data = []int{1, 2, 3, 4, 5}
	for _, nilai := range data {
		fmt.Println("Nilai:", nilai)
	}
}
