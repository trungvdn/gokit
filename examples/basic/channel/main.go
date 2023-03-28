package main

import "fmt"

func main() {
	fmt.Println("Main start")
	stringChan := make(chan string)

	go func() {
		stringChan <- "message"
	}()

	fmt.Println(<-stringChan)
	fmt.Println("Main finish")
}
