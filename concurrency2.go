package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Used to wait for the program to finish
var wg sync.WaitGroup

func main() {

	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// Allocate a logical processor for every available core.
	//runtime.GOMAXPROCS(runtime.NumCPU())


	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")

	wg.Wait()

	fmt.Println("\n...Program exiting...")
}

func printPrime(prefix string) {
	defer wg.Done()

	next:
		for outer := 2; outer < 5000; outer++ {
			for inner := 2; inner < outer; inner++ {
				if outer%inner == 0 {
					continue next
				}
			}
			fmt.Printf("%s:%d\n", prefix, outer)
		}
		fmt.Println("Completed", prefix)
}