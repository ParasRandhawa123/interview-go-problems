package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type job struct {
	index int
	value int
}

type result struct {
	index int
	value int
	err   error
}

func worker(ctx context.Context, jobs <-chan job, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		select {
		case <-ctx.Done():
			return
		default:
		}
		results <- result{index: j.index, value: j.value * j.value}
	}
}

func processJobs(ctx context.Context, input []int, numWorkers int) ([]int, error) {
	jobs := make(chan job, len(input))
	results := make(chan result, len(input))

	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, jobs, results, &wg)
	}

	for i, v := range input {
		jobs <- job{index: i, value: v}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	out := make([]int, len(input))
	received := 0
	for received < len(input) {
		select {
		case r, ok := <-results:
			if !ok {
				return out, fmt.Errorf("cancelled or workers exited early: got %d/%d results", received, len(input))
			}
			if r.err != nil {
				return out, r.err
			}
			out[r.index] = r.value
			received++
		case <-ctx.Done():
			return out, ctx.Err()
		}
	}
	return out, nil
}

func main() {
	input := make([]int, 10)
	for i := range input {
		input[i] = i
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	out, err := processJobs(ctx, input, 3)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(out)
}
