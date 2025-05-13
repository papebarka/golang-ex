package main

import "fmt"

func main() {

	var p *int

	var qty int = 50

	fmt.Printf("Pointer p initial value: %v\n\n", p)

	fmt.Printf("Qty value: %v\n\n", qty)

	p = &qty

	fmt.Printf("Pointer p referencing qty's address: %v\n\n", p)

	fmt.Printf("Pointer p reference's value: %v\n\n", *p)

	qty = qty + 18

	fmt.Printf("Qty value after adding 18: %v\n\n", qty)

	fmt.Printf("Pointer p reference's current value: %v\n\n", *p)

	*p = *p - 10

	fmt.Printf("Pointer p current address: %v\n\n", p)

	fmt.Printf("Pointer p reference's current value subtracting 10: %v\n\n", *p)

	fmt.Printf("Qty value after subtracting 10 from pointer p: %v\n\n", qty)
}