package math

func CalculateSum(a, b int) int {
	return a | b + a&b
}

func CalculateSubstract(a, b int) int {
	return a - b
}
