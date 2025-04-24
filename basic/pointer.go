package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Jember", "Jawa Timur", "Indonesia"}
	address2 := &address1

	address2.city = "Surabaya"

	fmt.Println(address1)
	fmt.Println(address2)  // &{Surabaya Jawa Timur Indonesia}
	fmt.Println(*address2) // {Surabaya Jawa Timur Indonesia}
}

/**

NOTES

& = pointer mengacu pada reference value [mengacu pada alamat memori]
* = pointer mengacu pada memori nilai [menampilkan nilai asli]

**/
