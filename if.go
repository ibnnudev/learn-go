package main

import "fmt"

func main() {
	var missingNumber int
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		if number == 5 {
			continue
		}
		if number == 8 {
			missingNumber = number
			break
		}
	}
	fmt.Println("Missing number:", missingNumber)
	// Output: Missing number: 8
}