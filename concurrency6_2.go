// Buffered channels

/*
 *
 * A buffered channel is a channel with capacity to hold one or more values before they’re
 * received. These types of channels don’t force goroutines to be ready at the same
 * instant to perform sends and receives. There are also different conditions for when a
 * send or receive does block. A receive will block only if there’s no value in the channel
 * to receive. A send will block only if there’s no available buffer to place the value being
 * sent. This leads to the one big difference between unbuffered and buffered channels:
 * An unbuffered channel provides a guarantee that an exchange between two goroutines
 * is performed at the instant the send and receive take place. A buffered channel
 * has no such guarantee.
 *
 */


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4	// Number of goroutines to use.
	taskLoad		 = 10	// Amount of work to process.
)
// Used to wait for the program to finish
var wg sync.WaitGroup


func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	//Create a buffered channel to manage the task load
	tasks := make(chan string, taskLoad)

	// Launch goroutines to handle the work
	wg.Add(numberGoroutines)
	for gr:= 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// Add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// Close the channel so the goroutines will quit
	// when all the work is done.
	close(tasks)

	wg.Wait()

}


// Player simulates a person playing the game of tennis.

func worker(tasks chan string, worker int) {
	// Report that we just returned.
	defer wg.Done()

	for {

		// Wait for work to be assigned.
		task, ok := <-tasks
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Worker %d : Shutting down\n", worker)
			return
		}

		// Display we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finised the work.
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)

	}
	
}