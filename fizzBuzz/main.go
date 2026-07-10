package main

import (
	"fmt"
	"sync"
)

func fizz(ch chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		fmt.Println("Fizz")
		done <- struct{}{}
	}
}

func buzz(ch chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		fmt.Println("Buzz")
		done <- struct{}{}
	}
}

func fizzBuzz(ch chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		fmt.Println("FizzBuzz")
		done <- struct{}{}
	}
}

func numbers(ch chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Println(n)
		done <- struct{}{}
	}
}

func main() {
	n := 16
	var wg sync.WaitGroup

	fizzChan := make(chan int)
	buzzChan := make(chan int)
	fizzBuzzChan := make(chan int)
	numChan := make(chan int)
	done := make(chan struct{})

	go fizz(fizzChan, done, &wg)
	go buzz(buzzChan, done, &wg)
	go fizzBuzz(fizzBuzzChan, done, &wg)
	go numbers(numChan, done, &wg)

	wg.Add(4)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fizzBuzzChan <- i
		} else if i%3 == 0 {
			fizzChan <- i
		} else if i%5 == 0 {
			buzzChan <- i
		} else {
			numChan <- i
		}
		<-done
	}

	close(fizzChan)
	close(buzzChan)
	close(fizzBuzzChan)
	close(numChan)

	wg.Wait()
}
