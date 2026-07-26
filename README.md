# Go Concurrency Interview Problems

Solutions to Go concurrency interview practice questions — goroutines, channels,
synchronization, and `context`. Practice prompts are in
[go_concurrency_interview_questions.pdf](go_concurrency_interview_questions.pdf).

## Solutions

| Folder                                                           | Question                                           | Description                                                                                                                                                                                                                         |
| ---------------------------------------------------------------- | -------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [oddEven](oddEven/main.go)                                       | Q1 — Print Even and Odd Numbers Alternately        | Two goroutines ping-pong via two channels to print 0..N in order.                                                                                                                                                                   |
| [ABC_goroutine](ABC_goroutine/main.go)                           | Q2 — Three Goroutines Printing A, B, C in Sequence | Three goroutines pass a token through channels to print `ABCABC...`.                                                                                                                                                                |
| [pingPong](pingPong/main.go)                                     | Q3 — Ping-Pong Between Two Goroutines              | Two goroutines alternate printing `Ping`/`Pong` N times over two channels.                                                                                                                                                          |
| [1-to-N](1-to-N/main.go)                                         | Q4 — Print 1 to N Using K Goroutines in Order      | K goroutines print 1..N in round-robin order using a ring of channels.                                                                                                                                                              |
| [fizzBuzz](fizzBuzz/main.go)                                     | Q5 — Concurrent FizzBuzz with Four Goroutines      | Four goroutines (Fizz/Buzz/FizzBuzz/Number) each handle their category, coordinated via a shared `done` channel.                                                                                                                    |
| [singleProducerConsumer](singleProducerConsumer/main.go)         | Q6 — Bounded Buffer, Single Producer/Consumer      | One producer sends 1..N on a buffered channel; one consumer prints squares.                                                                                                                                                         |
| [multiProduceSingleConsumer](multiProduceSingleConsumer/main.go) | Q7 — Multiple Producers, Single Consumer           | P producer goroutines feed one channel; a `sync.WaitGroup` signals when to close it so a single consumer can sum the total ([main_v2.go](multiProduceSingleConsumer/main_v2.go) adds an explicit consumer goroutine + `WaitGroup`). |
| [workerPool](workerPool/main.go)                                 | Q8 / Q23 — Worker Pool (+ cancellation)            | A `context`-aware worker pool of W workers processes jobs concurrently, collects results in input order, and propagates errors or cancellation. Covered by [worker_test.go](workerPool/worker_test.go).                             |
| [fanInOut](fanInOut)                                             | Q11 / Q12 — Fan-Out / Fan-In                       | Placeholder, not yet implemented.                                                                                                                                                                                                   |

## Running

Each folder is an independent `main` package.

```bash
go run ./oddEven
go run ./ABC_goroutine
go run ./pingPong
go run ./1-to-N
go run ./fizzBuzz
go run ./singleProducerConsumer
go run ./multiProduceSingleConsumer
go run ./workerPool
```

Run with the race detector, as the PDF recommends:

```bash
go run -race ./workerPool
```

Run tests:

```bash
go test ./...
```

## Reference

[go_concurrency_interview_questions.pdf](go_concurrency_interview_questions.pdf) contains 35 practice
questions across 8 sections: Coordination & Alternation, Producer-Consumer Patterns, Fan-Out/Fan-In,
Synchronization Primitives, Channel Patterns & Pitfalls, Context & Cancellation, Real-World Scenarios,
and Gotchas Interviewers Love. Only a subset are solved here so far (see table above).
