package helper

import "fmt"

func StringToNumber(str string) (int, error) {
	var num int
	_, err := fmt.Sscanf(str, "%d", &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func StringToFloat(str string) (float64, error) {
	var num float64
	_, err := fmt.Sscanf(str, "%f", &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}
