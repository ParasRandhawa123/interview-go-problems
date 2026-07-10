package main

import (
	"fmt"
	"sync"
)

func A(num int, chanA chan struct{}, chanB chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		<-chanA
		fmt.Print("A")
		chanB <- struct{}{}
	}
}

func B(num int, chanB chan struct{}, chanC chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		<-chanB
		fmt.Print("B")
		chanC <- struct{}{}
	}
}

func C(num int, chanC chan struct{}, chanA chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < num; i++ {
		<-chanC
		fmt.Print("C")
		if i < num-1 {
			chanA <- struct{}{}
		}
	}
}

// ABCABCABC
func main() {
	num := 5
	chanA := make(chan struct{})
	chanB := make(chan struct{})
	chanC := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(3)
	go A(num, chanA, chanB, &wg)
	go B(num, chanB, chanC, &wg)
	go C(num, chanC, chanA, &wg)

	chanA <- struct{}{}

	wg.Wait()
	fmt.Println()
}
