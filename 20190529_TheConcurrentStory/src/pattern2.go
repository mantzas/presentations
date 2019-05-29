package main

import (
	"fmt"
	"sync"
)

// GEN START OMIT
func gen1(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// GEN END OMIT

// SQUARE START OMIT
func sq1(v int, chRes chan<- int) {
	chRes <- v * v
}

// SQUARE END OMIT

// WORK START OMIT
func distr(ch <-chan int, chRes chan<- int) {
	wg := sync.WaitGroup{}
	for v := range ch {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			sq1(v, chRes)
		}(v)
	}
	wg.Wait()
	close(chRes)
}

// WORK END OMIT

func main() {
	// MAIN START OMIT
	numbers := make(chan int, 10)
	results := make(chan int, 10)
	go gen1(numbers)
	go distr(numbers, results)
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Print("Finished\n")
	// MAIN END OMIT
}
