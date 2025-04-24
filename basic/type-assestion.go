package main

import "fmt"

func random() any {
	return "200 OK"
}

func main() {
	result := random()

	// !traditional way
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// !safe way
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Boolean", value)
	default:
		fmt.Println("type unknown")
	}
}
