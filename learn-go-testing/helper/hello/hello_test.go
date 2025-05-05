package hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSayHello(t *testing.T) {
	result := SayHello()
	expected := "Hello, World!"
	require.Equal(t, expected, result, "The SayHello function did not return the expected result")
}

func TestSayHelloTo(t *testing.T) {
	name := "Alice"
	result := SayHelloTo(name)
	expected := "Hello, Alice!"
	require.Equal(t, expected, result, "The SayHelloTo function did not return the expected result")

	name = "Bob"
	result = SayHelloTo(name)
	expected = "Hello, Bob!"
	require.Equal(t, expected, result, "The SayHelloTo function did not return the expected result for Bob")
}

func BenchmarkSayHello(b *testing.B) {
	for b.Loop() {
		SayHello()
	}
}
