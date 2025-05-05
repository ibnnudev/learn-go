package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestWithValue(t *testing.T) {
	type contextKey string

	contextA := context.Background()

	contextB := context.WithValue(contextA, contextKey("b"), "B")
	contextC := context.WithValue(contextA, contextKey("c"), "C")

	contextD := context.WithValue(contextB, contextKey("d"), "E")
	contextE := context.WithValue(contextB, contextKey("e"), "E")

	contextF := context.WithValue(contextC, contextKey("f"), "F")
	contextG := context.WithValue(contextF, contextKey("g"), "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println("--------------------")

	fmt.Println(contextA.Value(contextKey("a")))
	fmt.Println(contextB.Value(contextKey("b")))
	fmt.Println(contextC.Value(contextKey("c")))
	fmt.Println(contextD.Value(contextKey("d")))
	fmt.Println(contextE.Value(contextKey("e")))
	fmt.Println(contextF.Value(contextKey("f")))
	fmt.Println(contextG.Value(contextKey("g")))
}

func CreateCounter(context context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-context.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total goroutines:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Total goroutines:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total goroutines:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)

	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutines:", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total goroutines:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(6*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total goroutines:", runtime.NumGoroutine())
}
