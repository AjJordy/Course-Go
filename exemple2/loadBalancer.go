package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)

	qtdWorkers := 10000
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, canal) // go routine
	}
	// go worker(1, canal) // go routine
	// go worker(2, canal) // go routine
	// go worker(3, canal) // go routine

	for i := 0; i < 10000; i++ {
		canal <- i
	}
}
