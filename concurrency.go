package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		fmt.Println("Printing digits")

		for count := 0; count < 3; count++ {
			for num := 1; num < 27; num++ {
				fmt.Printf("%d", num)
			}
		}
	}()

	go func() {
		defer wg.Done()

		fmt.Println("Printing characters")

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting routines to finish")

	wg.Wait()

	fmt.Println("\n...Program exiting...")
}
