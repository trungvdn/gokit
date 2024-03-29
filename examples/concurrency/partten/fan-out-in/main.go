package main

import (
	"fmt"
	"sync"
)

// Multiple functions can read from the same channel until that channel is closed
func main() {
	in := gen(2, 3)

	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

}
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// func merge(cs ...<-chan int) <-chan int {
// 	out := make(chan int)
// 	var wg sync.WaitGroup
// 	for _, c := range cs {
// 		wg.Add(1)
// 		go func(c <-chan int) {
// 			for n := range c {
// 				out <- n
// 			}
// 			wg.Done()
// 		}(c)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }
