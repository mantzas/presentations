package main

import (
	"fmt"
	"sync"
)

// TYPE START OMIT
type sum struct {
	sync.Mutex
	sum int
}

func (s *sum) add(i int) {
	s.Lock()
	defer s.Unlock()
	s.sum = s.sum + 1
}

// TYPE END OMIT

func main() {
	// MAIN START OMIT
	wg := sync.WaitGroup{}
	s := sum{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.add(i)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Sum: %d\n", s.sum)
	// MAIN END OMIT
}
