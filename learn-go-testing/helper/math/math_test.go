package math

import (
	"fmt"
	"testing"
)

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{1, 2, 3},
	{2, 3, 5},
	{3, 4, 7},
	{4, 5, 9},
	{5, 6, 11},
}

func TestCalculateSum(t *testing.T) {
	for _, test := range addTests {
		if output := CalculateSum(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

type subtractTest struct {
	arg1, arg2, expected int
}

var subtractTests = []subtractTest{
	{2, 1, 1},
	{2, 2, 0},
	{5, 3, 2},
}

func TestCalculateSubtract(t *testing.T) {
	for _, st := range subtractTests {
		if output := CalculateSubstract(st.arg1, st.arg2); output != st.expected {
			t.Errorf("Output %q not equal to expected %q", output, st.expected)
		}
	}
}

func BenchmarkCalculateSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSum(4, 5)
	}
}

func ExampleCalculateSum() {
	fmt.Println(CalculateSum(4, 5))
	// Output: 9
}
