package main

import "fmt"

type Student struct {
	Name, Class string
	Grade       int
}

func (s Student) sayHello() {
	fmt.Printf("Hello, my name is %s, I am in class %s and my grade is %d\n", s.Name, s.Class, s.Grade)
}

func (s *Student) changeName(newName string) {
	s.Name = newName
}

func main() {
	student1 := Student{
		Name:  "John Doe",
		Class: "10A",
		Grade: 90,
	}

	student1.sayHello()
	student1.changeName("Jane Doe")
	student1.sayHello()
}
