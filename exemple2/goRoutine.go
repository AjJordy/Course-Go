package main

import (
	"fmt"
	"time"
)

func contador(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	go contador(10) // Go routine
	go contador(10) // Go routine
	contador(10)
}
