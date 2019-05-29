package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan string)

	go func() {
		fmt.Print(<-ch)
		close(ch)
	}()

	ch <- "Hello"
	<-ch
	fmt.Print(", World!\n")
	// END OMIT
}
