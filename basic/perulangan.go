package main

import "fmt"

func main() {
	// Perulangan dengan for
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	// Perulangan dengan while
	j := 0
	for j < 5 {
		fmt.Println("Perulangan ke-", j)
		j++
	}

	// Perulangan tanpa batas
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("Perulangan ke-", k)
		k++
	}

	// Perulangan dengan range
	slice := []string{"apel", "jeruk", "pisang"}
	for index, value := range slice {
		fmt.Println("Index:", index, "Value:", value)
	}

	// case : piramida *
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// case : piramida terbalik *
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// case : piramida angka
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
