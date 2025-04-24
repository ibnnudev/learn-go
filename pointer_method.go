package main

import "fmt"

type Teacher struct {
	Name string
}

func (p *Teacher) setName() {
	p.Name = "Nama saya " + p.Name
}

func main() {
	ibnu := Teacher{"Ibnu"}
	ibnu.setName()

	fmt.Println(ibnu.Name)
}
