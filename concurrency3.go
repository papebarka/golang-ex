package main

import (
	"fmt"
	"runtime"
	"sync"
)


var (
	// Variable incremented by all goroutines
	counter int
	// Used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {

	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	// Allocate a logical processor for every available core.
	//runtime.GOMAXPROCS(runtime.NumCPU())


	wg.Add(2)

	//fmt.Println("Create Goroutines")

	go incCounter(1)
	go incCounter(2)

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capure the value of Counter
		value := counter

		// Yield the thread and be placed backinqueue
		runtime.Gosched()

		value++

		// Store the value back into Counter
		counter = value
	}
}


// Go has a special tool that can detect race conditions in your code. It’s extremely
// useful to find these types of bugs, especially when they’re not as obvious as our example.

//go build -race