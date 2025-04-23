package main

import "fmt"

func main() {
	var cheese = map[string]string{
		"cheddar": "yellow",
		"brie":    "white",
		"gouda":   "orange",
	}

	for k, v := range cheese {
		fmt.Println(k, v)
	}

	// delete item map
	delete(cheese, "brie")
	fmt.Println("After deletion:", cheese)

	cheese["feta"] = "white"
	fmt.Println("After adding feta:", cheese)

	// check item exist
	var value, isExist = cheese["cheddars"]

	if isExist {
		fmt.Println((value))
	} else {
		fmt.Println("item is not exist")
	}

	var animals = []map[string]string{
		{"nama": "kangguru", "gender": "male"},
		{"nama": "kucing", "warna": "hitam"},
		{"nama": "anjing", "warna": "putih"},
	}

	for _, animal := range animals {
		fmt.Println(animal)
	}
}
