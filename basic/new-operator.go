package main

import "fmt"

type Person struct {
	name, gender string
}

func main() {
	var person1 *Person = new(Person)
	var person2 *Person = person1

	person2.name = "Hendy"

	fmt.Println(person1)
	fmt.Println(person2)
}
