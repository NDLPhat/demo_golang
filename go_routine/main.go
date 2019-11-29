package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go Task1(messages)
	go Task2(messages)
	go Task3(messages)

	msg1 := <-messages
	msg2 := <-messages
	msg3 := <-messages
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)

}

func Task1(ch chan string) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task 1 Done")
	ch <- "task1"

}

func Task2(ch chan string) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task 2 Done")
	ch <- "task2"
}

func Task3(ch chan string) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task 3 Done")
	ch <- "task3"
}
