package main

import "fmt"

type AnimalInteface interface {
	details()
}

func details(a AnimalInteface) {
	fmt.Println("Animal details")
	a.details()
}

type AnimalType struct {
	Name    string
	Species string
}

func (a AnimalType) details() {
	fmt.Printf("Name: %s, Species: %s\n", a.Name, a.Species)
}

func main() {
	animal := AnimalType{
		Name:    "Lion",
		Species: "Panthera leo",
	}

	details(animal)
}
