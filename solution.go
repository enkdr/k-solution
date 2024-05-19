package main

import (
	"fmt"
	"sync"
)

// added waitgroup
// removed id
// limited number of workers

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // defer to wait for all workers to finish
	for j := range jobs {
		switch j % 3 {
		case 0:
			// j = j * 1 // unessessary operation
		case 1:
			j = j * 2
			results <- j * 2
		case 2:
			results <- j * 3
			j = j * 3
		}
	}
}

func main() {
	jobs := make(chan int)    // create channel to send jobs to workers
	results := make(chan int) // create channel to receive results from workers
	var wg sync.WaitGroup     // create waitgroup to wait for all workers to finish

	// limit number of workers -- can be increased depending on resources
	numWorkers := 100
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(jobs, results, &wg)
	}

	// moved loop inside
	go func() {
		for i := 1; i <= 1000000000; i++ {
			if i%2 != 0 {
				jobs <- i
			}
		}
		close(jobs)
	}()

	// wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// sum up all results
	var sum int
	for r := range results {
		sum += r
	}
	fmt.Println(sum)
}
