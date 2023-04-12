package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Race(chanel, quit chan string, i int) {
	chanel <- fmt.Sprintf("Car %d started", i)
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
	quit <- fmt.Sprintf("Car %d reached the finishing line!", i)
}

func main() {

	chanel := make(chan string)
	quit := make(chan string)
	for i := 0; i < 3; i++ {
		go Race(chanel, quit, i)
	}

	for {
		select {
		case startCar := <-chanel:
			fmt.Println(startCar)
		case endCar := <-quit:
			fmt.Println(endCar)
			return
		}
	}
}
