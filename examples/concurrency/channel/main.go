package main

import (
	"fmt"
	"time"
)

func square(c chan int) {
	fmt.Println("[square] reading")
	num := <-c
	c <- num * num
	fmt.Println("[square] done")
}

func cube(c chan int) {
	fmt.Println("[cube] reading")
	num := <-c
	c <- num * num * num
	fmt.Println("[cube] done")
}

func main() {
	fmt.Println("[main] main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	fmt.Println("[main] sent testNum to squareChan")

	squareChan <- testNum

	time.Sleep(time.Second)

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")

	cubeChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of square and cube of", testNum, " is", sum)
	fmt.Println("[main] main() stopped")
}
