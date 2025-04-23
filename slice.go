package main

import "fmt"

func main() {
	var students = []string{"ibnu", "jhonny", "rumi"}
	fmt.Println(students[0:2])

	// len
	fmt.Println("Panjang slice:", len(students))

	// append
	var newStudents = append(students, "budi", "andi")
	fmt.Println("Murid baru:", newStudents)

	// copy
	var copyStudents = make([]string, len(newStudents))
	copy(copyStudents, newStudents)
	fmt.Println("Copy students:", copyStudents)
}
