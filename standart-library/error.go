package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("error in validation")
	NotFoundError   = errors.New("error bcs not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "ibnu" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("budi")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println(ValidationError.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println(NotFoundError.Error())
		} else {
			fmt.Println("Unknown error")
		}
	}
}
