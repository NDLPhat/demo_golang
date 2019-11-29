package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	go func() {
		ch <- 200
	}()

	fmt.Println("hello world")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
