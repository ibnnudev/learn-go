package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	fmt.Println("2 pangkat 3 =", pangkat(2, 3))
	fmt.Println("Random number = ", randomizer.Intn(100))

	sayHello("ibnu")
}

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	} else if y == 1 {
		return x
	} else {
		return x * pangkat(x, y-1)
	}
}

func sayHello(name string) {
	fmt.Println("halo", name)
}
