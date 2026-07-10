package main

import (
	"fmt"
	"sync"
)

func producer(n int, ch chan int) {
	defer close(ch)
	for i := 1; i <= n; i++ {
		ch <- i
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i * i)
	}
}

func main() {
	ch := make(chan int, 5)
	n := 10

	var wg sync.WaitGroup

	wg.Add(1)

	go producer(n, ch)
	go consumer(ch, &wg)

	wg.Wait()
}
