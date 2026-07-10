package main

import (
	"fmt"
	"sync"
)

func oddFunc(num int, oddChan chan struct{}, evenChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= num; i += 2 {
		<-oddChan
		fmt.Println(i)
		if i < num {
			evenChan <- struct{}{}
		}
	}
}

func evenFunc(num int, evenChan chan struct{}, oddChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= num; i += 2 {
		<-evenChan
		fmt.Println(i)
		if i < num {
			oddChan <- struct{}{}
		}
	}
}

func main() {
	num := 99
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)

	go oddFunc(num, oddChan, evenChan, &wg)

	go evenFunc(num, evenChan, oddChan, &wg)

	evenChan <- struct{}{}
	wg.Wait()
}
