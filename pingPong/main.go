package main

import (
	"fmt"
	"sync"
)

func ping(num int, pingCh chan struct{}, pongCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		<-pingCh
		fmt.Println("Ping")
		pongCh <- struct{}{}
	}
}

func pong(num int, pongCh chan struct{}, pingCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= num; i++ {
		<-pongCh
		fmt.Println("Pong")
		if i != num {
			// till "i" is smaller than num, as when its equal no pingCh reviever is there if read from channel
			pingCh <- struct{}{}
		}
	}
}

func main() {
	num := 2
	pingChan := make(chan struct{})
	pongChan := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go ping(num, pingChan, pongChan, &wg)
	go pong(num, pongChan, pingChan, &wg)

	pingChan <- struct{}{}
	wg.Wait()

	fmt.Println("Game Done !!")
}
