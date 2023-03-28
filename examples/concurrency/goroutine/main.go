package main

import (
	"fmt"
	"time"
)

func main() {

	// var wg sync.WaitGroup
	startTime := time.Now()
	for i := 0; i < 5; i++ {
		// wg.Add(1)
		i := i
		func() {
			// defer wg.Done()
			time.Sleep(2 * time.Second)
			fmt.Printf("done with %d\n", i)
		}()
	}
	// wg.Wait()
	endTime := time.Since(startTime)
	fmt.Println("Total: ", endTime.Milliseconds())
}
