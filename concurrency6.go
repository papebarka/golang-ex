/*
 * Atomic functions and mutexes work, but they don’t make writing concurrent programs
 * easier, less error-prone, or fun. In Go you don’t have only atomic functions and
 * mutexes to keep shared resources safe and eliminate race conditions. You also have
 * channels that synchronize goroutines as they send and receive the resources they
 * need to share between each other
 *
 *
 * When a resource needs to be shared between goroutines, channels act as a conduit
 * between the goroutines and provide a mechanism that guarantees a synchronous
 * exchange. When declaring a channel, the type of data that will be shared needs to be
 * specified. Values and pointers of built-in, named, struct, and reference types can be
 * shared through a channel.
 *
 */

 /*
  * Creating a channel in Go requires the use of the built-in function make.
  * We can create both buffered or unbuffered channels.
  * The first argument to make requires the keyword chan and then the type of data the channel will allow to be exchanged.
  * If you’re creating a buffered channel, then you specify the size of the channel’s buffer as the second argument.
  * Sending a value or pointer into a channel requires the use of the <- operator
  */

  // buffered := make(chan string, 10) - Buffered channel
  // unbuffered := make(chan int) - Unbuffered channel
  // Send a string thourgh the channel
  // buffered <-"Gopher"
  // Receive a string from the channel
  // value := <-buffered


/* Unbuffered and buffered channels behave a bit differently. Understanding the differences
 * will help you determine when to prefer one over the other, so let’s look at
 * each type separately
 */

