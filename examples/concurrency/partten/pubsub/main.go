package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	messageChan := make(chan int, 10)
	maxParam := 50

	var wg sync.WaitGroup
	startTime := time.Now()
	// comsumer
	numWorker := 10
	worker := newWorker(messageChan, &wg)
	for i := 0; i < numWorker; i++ {
		go worker.start(i)
	}

	// publisher
	for i := 0; i < maxParam; i++ {
		fmt.Printf("Producing with param %d\n", i)
		wg.Add(1)
		messageChan <- i
	}
	endTime := time.Since(startTime)
	fmt.Println("Done all job in : ", endTime.Milliseconds())
	wg.Wait()
	fmt.Println("num of goroutine after execute all job :", runtime.NumGoroutine())
	//TO DO: close goroutine un-use
}

type worker struct {
	queue chan int
	wg    *sync.WaitGroup
}

func newWorker(queue chan int, wg *sync.WaitGroup) *worker {
	return &worker{
		queue: queue,
		wg:    wg,
	}
}

func (w *worker) start(id int) {
	fmt.Printf("Init goroutine %d to execute job\n", id)
	for {
		msg, ok := <-w.queue
		if ok {
			fmt.Println("channel has closed")
			break
		}
		w.excuteJob(msg)
	}
}

func (w *worker) excuteJob(param int) {
	defer w.wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("job have finished with param: %d\n", param)
}
