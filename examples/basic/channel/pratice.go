package main

import (
	"fmt"
	"time"
)

func square(c chan int) {
	fmt.Println("[square] reading and write to chanel")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("[cube] reading and write to chanel")
	num := <-c
	c <- num * num * num
}

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 2"
}

func ma2in() {

	fmt.Println("[main] started at ", time.Since(start))
	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service1", res)
	case res := <-chan2:
		fmt.Println("Response from service2", res)
	case <-time.After(2 * time.Second):
		fmt.Println("time out")
	}
	fmt.Println("[main] stopped at ", time.Since(start))
}
