package main

import "fmt"

func fib(length int) <-chan int {
	c := make(chan int, 10)
	go func() {
		for i, j := 0, 1; i < length; i, j = i+j, i {
			c <- i
		}
		close(c)

	}()
	return c
}

func main() {
	for f := range fib(10) {
		fmt.Println(f)
	}
}
