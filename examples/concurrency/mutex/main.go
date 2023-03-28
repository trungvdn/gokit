package main

import (
	"fmt"
	"sync"
)

var (
	count = 0
	wg    = &sync.WaitGroup{}
	mutex = &sync.Mutex{}
)

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()
			mutex.Lock()
			count++
		}()
	}
	wg.Wait()
	fmt.Println("count", count)

}
