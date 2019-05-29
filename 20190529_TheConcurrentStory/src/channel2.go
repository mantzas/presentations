package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan int)
	val, ok := <-ch
	fmt.Printf("val: %d ok: %t\n", val, ok)
	// END OMIT
}
