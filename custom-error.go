package main

import "errors"

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound   = errors.New("data not found")
	ErrUnexpected = errors.New("unexpected error")
)

func SaveData(id string, data any) error {
	if id == "" {
		return ErrValidation
	}

	if id != "ibnu" {
		return ErrNotFound
	}

	if id == "andi" {
		return ErrUnexpected
	}

	return nil
}

func main() {
	var students = map[string]string{
		"ibnu": "nama saya ibnu",
		"budi": "aku budi",
		"andi": "aku andi",
	}

	for id, name := range students {
		err := SaveData(id, name)
		if err != nil {
			println(err.Error(), id, name)
		} else {
			println("Data saved successfully", id, name)
		}
	}
}
