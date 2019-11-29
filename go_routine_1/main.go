package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Task1(&wg)
	go Task2(&wg)
	Task3()
	wg.Wait()

}

func Task1(wg *sync.WaitGroup) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task 1 Done")
	wg.Done()

}

func Task2(wg *sync.WaitGroup) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task 2 Done")
	wg.Done()
}

func Task3() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Task 3 Done")
	// ch <- "task3"
}
