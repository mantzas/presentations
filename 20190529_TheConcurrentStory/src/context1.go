package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// SETUP START OMIT
	chDone := make(chan struct{})
	ch := make(chan int, 1000)
	ctx, cnl := context.WithCancel(context.Background())
	// SETUP END OMIT

	// PRODUCER START OMIT
	go func() {
		i := 0
		for {
			if ctx.Err() != nil {
				fmt.Print("Producer Done\n")
				chDone <- struct{}{}
				return
			}
			ch <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()
	// PRODUCER END OMIT

	// CONSUMER START OMIT
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Print("Writer Done\n")
				chDone <- struct{}{}
				return
			case n := <-ch:
				fmt.Printf("%d\n", n)
			}
		}
	}()
	// CONSUMER END OMIT

	// CLOSING START OMIT
	time.Sleep(10 * time.Second)
	cnl()
	<-chDone
	<-chDone
	// CLOSING END OMIT
	fmt.Print("Finish\n")
}
