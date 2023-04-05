package main

import (
	"fmt"
	"sync"
	"time"
)

func sqrWorker(wg *sync.WaitGroup, task <-chan int, results chan<- int, instance int) {
	for num := range task {
		time.Sleep(time.Millisecond)
		fmt.Printf("[sqrWorker %v] Sending by worker %v\n", instance, instance)
		results <- num * num
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("[main] main started")
	task := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(&wg, task, results, i)
	}

	for i := 0; i < 5; i++ {
		task <- i * 2
	}

	fmt.Println("[main] wrote 5 tasks")
	close(task)
	wg.Wait()

	// for res := range results {
	// 	// res := <-results
	// 	fmt.Println("[main] result = ", res)
	// }
	for i := 0; i < 5; i++ {
		res := <-results
		fmt.Println("[main] result = ", res)

	}
	fmt.Println("[main] main closed")
}
