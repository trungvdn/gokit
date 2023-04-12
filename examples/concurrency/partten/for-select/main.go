package main

import "fmt"

func getChampion(c chan string) {
	chams := []string{"Caitlyn", "Jhin", "Jinx", "Xaya"}
	for _, champion := range chams {
		c <- champion
	}
	c <- "Done"
	close(c)
}

func main() {
	chamChan := make(chan string)

	go getChampion(chamChan)

	for {
		select {
		case cham := <-chamChan:
			fmt.Println("cham =", cham)
			if cham == "Done" {
				fmt.Println("done ", cham)
				return
			}
		default:
		}
	}
}
