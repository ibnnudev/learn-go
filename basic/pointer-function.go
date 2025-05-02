package main

import "fmt"

type Address struct {
	Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address *Address = &Address{}
	changeCountryToIndonesia(address)

	fmt.Println(address)
}
