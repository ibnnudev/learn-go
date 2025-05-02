package main

import (
	database "learn-go/database"
	_ "learn-go/internal"
)

func main() {
	println(database.GetDatabase())
}
