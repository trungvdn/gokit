package main

import (
	"fmt"
)

func greeting(c chan string) {
	<-c
}

// func square(c chan int) {
// 	for i := 0; i < 9; i++ {
// 		c <- i + i
// 	}
// 	close(c)
// }

func ma2n() {
	// Declaring a channel
	var c1 chan string
	fmt.Printf("value of c,type %v,%T\n", c1, c1)
	c2 := make(chan string)
	fmt.Printf("value of c,type %v,%T\n", c2, c2)
	// Data read and write
	// data := <-c1
	// c2 <- "Hello"
	// fmt.Println(data)
	// Deadlock

	// Closing a channel

	// c3 := make(chan string)

	// go greeting(c3)
	// c3 <- "John"

	// close(c3)

	// c3 <- "Jce"

	// For loop
	// c4 := make(chan int)
	// go square(c4)

	// for {
	// 	val, ok := <-c4
	// 	if ok {
	// 		fmt.Println(val)
	// 	} else {
	// 		fmt.Println(val, ok)
	// 		break
	// 	}
	// }
	// time.Sleep(2 * time.Minute)

	// buffer channel

	c := make(chan int, 3)

	// go readChan(c)

	c <- 1
	c <- 2
	num := <-c
	fmt.Println(num)
	// fmt.Println("Main start")
	// stringChan := make(chan string)

	// go func() {
	// 	stringChan <- "message"
	// }()

	// fmt.Println(<-stringChan)
	fmt.Println("Main finish")

	// select
	//is just like switch without any input argument
}

func readChan(c chan int) {
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}
}
