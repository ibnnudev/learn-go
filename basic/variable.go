package main

import "fmt"

func main()  {
	var firstName string;
	var lastName string;

	firstName = "Ibnu";
	lastName = "Nugroho";

	fmt.Println("Hello " + firstName + " " + lastName);

	// without var
	age := 20;
	fmt.Println("Age: ", age);
}