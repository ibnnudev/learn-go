package main

import "fmt"

func main() {
	var missingNumber int = 9
	if missingNumber == 9 {
		fmt.Println("The number is 9")
	} else {
		fmt.Println("The number is not 9")
	}

	const name string = "ibnu"

	switch name {
	case "ibnu":
		fmt.Println("The name is ibnu")
	default:
		fmt.Println("The name is not ibnu")
		fmt.Println("This is the default case")
	}

	const minGradeToGraduate = 91
	if minGradeToGraduate > 90 {
		fmt.Println("You are graduated")
	} else if minGradeToGraduate == 90 || minGradeToGraduate < 90 {
		fmt.Println("You are not graduated")
	}

	const lastPoint int = 100
	if lastPoint > 90 {
		if lastPoint == 100 {
			fmt.Println("You got a perfect score")
		} else {
			fmt.Println("You got a score greater than 90")
		}
	} else {
		fmt.Println("You got a score of 90 or less")
	}
}
