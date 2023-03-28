package main

import "fmt"

// a pipeline is a series of stages connected by channels
func main() {
	c := gen(1, 2, 3, 4)
	for v := range sq(c) {
		fmt.Println(v)
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
