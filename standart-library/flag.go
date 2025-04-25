package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put ur database host")
	username := flag.String("username", "root", "Put ur database username")
	password := flag.String("password", "root", "Put ur database password")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
}
