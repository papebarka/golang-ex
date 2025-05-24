// Unbuffered channels

/*
 *
 * An unbuffered channel is a channel with no capacity to hold any value before it’s
 * received. These types of channels require both a sending and receiving goroutine to
 * be ready at the same instant before any send or receive operation can complete. If the
 * two goroutines aren’t ready at the same instant, the channel makes the goroutine that
 * performs its respective send or receive operation first wait. Synchronization is inherent
 * in the interaction between the send and receive on the channel. One can’t happen
 * without the other.
 *
 */


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Used to wait for the program to finish
var wg sync.WaitGroup


func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	//Create an unbuffered channel
	court := make(chan int)

	wg.Add(2)

	// Launch two players

	go player("Alassane", court)
	go player("Alousseyn", court)
	
	// Start the set
	court <- 1

	wg.Wait()

}


// Player simulates a person playing the game of tennis.

func player(name string, court chan int) {
	defer wg.Done()

	for {

		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player
		court <- ball
	}
	
}