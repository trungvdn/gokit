package main

import (
	"fmt"
	"time"
)

func main() {
	// ctx := context.Background()

	// ctx, cancel := context.WithCancel(ctx)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cancel()
	// }()

	// select {
	// case <-time.After(2 * time.Second):
	// 	fmt.Println("Hello")
	// case <-ctx.Done():
	// 	log.Fatalf(ctx.Err().Error())
	// }
	a := 100
	fmt.Println("hello")
	fmt.Println(a)
	fmt.Println(time.Now())
}
