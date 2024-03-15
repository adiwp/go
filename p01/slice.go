package main

import "fmt"

func main() {
	names := [...]string{"Adi", "Amir", "Masbro", "Greg", "Bambang", "Ion", "Desti", "Sri"}
	slice := names[2:6] // slice dari elemen index 2 hingga 5

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])

	// append
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(days)

	daySlace2 := append(daySlice1, "Libur baru")
	daySlace2[0] = "Ups"
	fmt.Println(daySlace2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fahri"
	newSlice[1] = "Darmawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	fmt.Println("fromSlice = ", fromSlice)
	fmt.Println("toSlice = ", toSlice)

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)
}
