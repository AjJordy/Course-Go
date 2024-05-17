package main

import (
	"fmt"
)

func main() {
	canal := make(chan string)
	go func() {
		canal <- "opa"
	}()

	msg := <-canal
	fmt.Println(msg)
}
