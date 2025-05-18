package main

import "fmt"

func Arrays() {
	var age [5]int

	fmt.Printf("Initialized array: %v\n", age)

	age[0], age[1], age[2], age[3], age[4] = 10, 7, 8, 9, 5

	fmt.Printf("Assigned array value: %v\n", age)

	name := [5]string{"John", "Dane", "Mike", "Damo", "Momo"}

	country := [...]string{"AU", "UK", "US", "ML", "CI"}

	fmt.Printf("Name: %v\n", name)

	fmt.Printf("Country: %v\n", country)
}

func Slices() {
	// Slice ages of length and capacity 5
	ages := make([]int, 5)

	// Slice names of length 3 and capacity 5
	names := make([]string, 3, 5)

	//Slice literal
	qty := []int{10, 82, 100, 35, 98, 14, 52}

	fmt.Println("************* SLICES *************\n")

	fmt.Printf("ages initial value: %v\n", ages)

	fmt.Printf("names initial value: %v\n", names)

	fmt.Printf("Qty initial value: %v\n", qty)

	// Assign a value to a specific slice index
	ages[0] = 18

	fmt.Printf("ages initialized value: %v\n", ages)

	// Adds value to a full slice capacity
	qty = append(qty, 288)
	fmt.Printf("Qty with appended value: %v\n", qty)

	for index, value := range qty {
		fmt.Printf("Qty at index %d is: %d\n", index, value)
	}

	/*
	* Passing slices to functions is trivial. Only the concerned slice is copied, not the underlining 	rray.
	* In 64-bit architecture, a slice requires 24  bytes of memory.
	* Pointer field -> 8 bytes
	* Length field -> 8 bytes
	* Capacity field -> 8 bytes
	*/
}

func main() {
	Arrays()

	Slices()
}