package main

import (
	"fmt"
)

// GEN START OMIT
func gen(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// GEN END OMIT

// SQUARE START OMIT
func sq(ch <-chan int, chRes chan<- int) {
	for v := range ch {
		chRes <- v * v
	}
	close(chRes)
}

// SQUARE END OMIT

func main() {
	// MAIN START OMIT
	numbers := make(chan int, 10)
	results := make(chan int, 10)
	go gen(numbers)
	go sq(numbers, results)

	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
	fmt.Print("Finished\n")
	// MAIN END OMIT
}
