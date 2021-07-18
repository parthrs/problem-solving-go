package main

import "fmt"

type car struct {
	color        string
	miles        int
	transmission string
}

func NewCar(color string, miles int, transmission string) *car {
	// Create a struct of type car in memory and return a pointer to it
	return &car{color, miles, transmission}
}

func main() {
	myC := NewCar("Red", 50, "Automatic")
	fmt.Printf("%v\n", myC)
}
