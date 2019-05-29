package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan string, 4)
	chSig := make(chan struct{})

	go func() {
		for val := range ch {
			fmt.Print(val)
		}
		chSig <- struct{}{}
	}()

	for i := 1; i < 4; i++ {
		ch <- fmt.Sprintf("%d,", i)
	}
	close(ch)
	<-chSig
	fmt.Print(" Boom!\n")
	// END OMIT
}
