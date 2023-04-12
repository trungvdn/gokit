package main

import "fmt"

type Money struct {
	amount int
	year   int
}

func sendMoney(m chan Money) {
	for i := 0; i < 18; i++ {
		m <- Money{
			amount: 5000,
			year:   i,
		}
	}
	close(m)
}

func main() {
	moneyChan := make(chan Money)

	go sendMoney(moneyChan)

	for m := range moneyChan {
		fmt.Printf("Money received by kid in year %d : %d\n", m.year, m.amount)
	}
}
