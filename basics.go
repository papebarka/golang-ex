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

	fmt.Println("************* SLICES *************\n")

	fmt.Printf("ages initial value: %v\n", ages)

	fmt.Printf("names initial value: %v\n", names)
}

func main() {
	Arrays()

	Slices()
}