package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ibnu Customer
	ibnu.Name = "Moh Ibnu A.S"
	ibnu.Address = "JKT"
	ibnu.Age = 23

	fmt.Println(ibnu)

	dela := Customer{
		Name:    "Dela A.S",
		Address: "Jember",
		Age:     26,
	}

	fmt.Println(dela)

	dela.Name = "Dela Sutio"
	fmt.Println("after edit:", dela)
}
