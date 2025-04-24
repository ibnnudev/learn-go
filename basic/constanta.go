package main

import "fmt"

func main() {
	const pi = 3.14
	const r = 4

	var luas = pi * r * r
	var keliling = 2 * pi * r
	fmt.Println("Luas Lingkaran = ", luas)
	fmt.Println("Keliling Lingkaran = ", keliling)
}