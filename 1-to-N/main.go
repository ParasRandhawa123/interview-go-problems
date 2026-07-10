package main

import (
	"fmt"
	"sync"
)

func worker(id, n, k int, in, out chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := id + 1; num <= n; num += k {
		<-in
		fmt.Println(num)
		if num+1 <= n {
			out <- struct{}{}
		}
	}
}

func main() {
	n := 10
	k := 3

	chans := make([]chan struct{}, k)
	for i := range chans {
		chans[i] = make(chan struct{})
	}

	var wg sync.WaitGroup
	wg.Add(k)

	for i := 0; i < k; i++ {
		go worker(i, n, k, chans[i], chans[(i+1)%k], &wg)
	}

	chans[0] <- struct{}{}
	wg.Wait()

}
