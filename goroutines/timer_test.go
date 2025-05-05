package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	time := <-timer.C
	fmt.Println(time)
}

func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	time := <-channel
	fmt.Println(time)
}

func TestTimeAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("5 seconds later")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		wg.Done()
	})
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	wg.Wait()
}
