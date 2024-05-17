package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func (c Car) Drive() {
	fmt.Println("The car is " + c.Name + " running")
}

func main() {
	println("Hello, world!")

	// a := 1 // Creating variable
	// a = 2  // Setting value to variable

	car1 := Car{Name: "Ford", Model: "Mustang", Year: 1969}
	car2 := Car{Name: "Chevrolet", Model: "Camaro", Year: 1969}

	fmt.Println(car1.Name)
	fmt.Println(car2.Name)

	car1.Drive()
	car2.Drive()

	// HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(car1)
	})
	http.ListenAndServe(":8080", nil)
}
