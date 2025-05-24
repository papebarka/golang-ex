package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Atomic functions provide low-level locking mechanisms for synchronizing access to
// integers and pointers. We can use atomic functions to fix race conditions

var (
	// Variable incremented by all goroutines
	counter int64
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
		// Safely Add One to Counter
		atomic.AddInt64(&counter, 1)

		// Yield the thread and be placed backinqueue
		runtime.Gosched()
	}
}


// Go has a special tool that can detect race conditions in your code. It’s extremely
// useful to find these types of bugs, especially when they’re not as obvious as our example.

//go build -race