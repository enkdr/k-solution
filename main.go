package main

import "fmt"

// problems

// Critical
// panic trying to send on closed channel

// Bad
// id is not being used in 'worker' function
// no wait group to handle closing channels
// in main the loop only send the numbers passing conditions - move loop

// Nice to have
// channels are not buffered so they will block

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // loop through jobs channel (the: if number is even then add 99 channel) for
		go func() {
			switch j % 3 {
			case 0: // j is divisible by 3?
				j = j * 1 // multiply j by 1 - unessessary operation
			case 1: // remainer is 1?
				j = j * 2 // multiply j by 2
				results <- j * 2
			case 2: // remainer is 2?
				results <- j * 3 // multiply j by 3
				j = j * 3
			}
		}()
	}
}

func main() {

	jobs := make(chan int)    // make an unbuffered channel
	results := make(chan int) // make an ubuffered channel

	// the loop here should be inside the go func
	for i := 1; i <= 1000000000; i++ { // loop a billion times
		go func() { // create a goroutine for each loop
			if i%2 == 0 { // if i is even
				i += 99 // add 99 to i
			} // eveytime i is even: 99 is added to i
			jobs <- i // send i to the channel
		}() // fun go func
	}
	close(jobs) // close the channel

	jobs2 := []int{} // create a new array of ints

	for w := 1; w < 1000; w++ { // loop 1000 times
		jobs2 = append(jobs2, w) // append w to the array
	}

	for i, w := range jobs2 { // loop through jobs2 array
		go worker(w, jobs, results) // call worker function on each job with (id int, jobs chan int, and results chan int)
		i = i + 1                   // increment i
	}

	close(results) // close the results channel

	var sum int32 = 0 // initialize sum to 0

	for r := range results { // loop through results channel
		sum += int32(r) // add r to sum
	}

	fmt.Println(sum) // print the sum
}
