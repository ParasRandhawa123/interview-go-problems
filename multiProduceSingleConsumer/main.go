package main

import (
	"fmt"
	"sync"
)

func main() {
	P := 100

	ch := make(chan int, P)

	var producerWg sync.WaitGroup

	for i := 1; i <= P; i++ {
		producerWg.Add(1)
		go func(val int) {
			defer producerWg.Done()
			ch <- val
		}(i)
	}

	go func() {
		producerWg.Wait()
		close(ch)
	}()

	sum := 0
	for v := range ch {
		sum += v
	}
	fmt.Println(sum)
}
