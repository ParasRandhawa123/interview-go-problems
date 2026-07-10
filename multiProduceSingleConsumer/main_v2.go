package main

import (
	"fmt"
	"sync"
)

func consumerV2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for i := range ch {
		result += i
	}
	fmt.Println(result)
}

func mainV2() {
	P := 100

	ch := make(chan int, P)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

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

	consumerWg.Add(1)

	go consumerV2(ch, &consumerWg)
	consumerWg.Wait()
}
