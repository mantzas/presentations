package main

import (
	"fmt"
	"sync"
)

// GEN START OMIT
func gen2(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// GEN END OMIT

// SQUARE START OMIT
func sq2(ch <-chan int, chRes chan<- int) {
	for v := range ch {
		chRes <- v * v
	}
}

// SQUARE END OMIT

// WORK START OMIT
func workers(ch <-chan int, chRes chan<- int) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sq2(ch, chRes)
		}()
	}
	wg.Wait()
	close(chRes)
}

// WORK END OMIT

func main() {
	// MAIN START OMIT
	numbers := make(chan int, 10)
	results := make(chan int, 10)
	go workers(numbers, results)
	go gen2(numbers)
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Print("Finished\n")
	// MAIN END OMIT
}
