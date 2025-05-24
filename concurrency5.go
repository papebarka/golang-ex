// This sample program demonstrates how to use a mutex
// to define critical sections of code that need synchronous access.


package main

import (
	"fmt"
	"runtime"
	"sync"
)

//Atomic functions provide low-level locking mechanisms for synchronizing access to
// integers and pointers. We can use atomic functions to fix race conditions
// Two other useful atomic functions are LoadInt64 and StoreInt64. These functions
// provide a safe way to read and write to an integer value

var (
	// Variable incremented by all goroutines
	counter int
	// Used to wait for the program to finish
	wg sync.WaitGroup
	// mutex is used to define a critical section of code.
	mutex sync.Mutex
)

func main() {

	wg.Add(2)

	//fmt.Println("Create Goroutines")

	go incCounter(1)
	go incCounter(2)

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Printf("Final Counter: %d\\n", counter)
}


// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.

		mutex.Lock()
		{
			// Capture the value of Counter
			value := counter

			// Yield the thread and be placed backinqueue
			runtime.Gosched()

			// Increment our local value of counter
			value++

			// Store the valye back into counter
			counter = value
		}
		mutex.Unlock()
		// Release the lock and allow any
		// waiting gorutine through.
	
	}
}


// Go has a special tool that can detect race conditions in your code. It’s extremely
// useful to find these types of bugs, especially when they’re not as obvious as our example.

//go build -race