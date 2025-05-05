package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	go func() {
		channel <- "Ibnu"
		channel <- "Sutio"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke:" + strconv.Itoa(i)
		}

		close(channel)
	}()

	for value := range <-channel {
		fmt.Println(value)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go func() {
		for _, i := range []int{0, 1} {
			channel1 <- "Data ke-" + strconv.Itoa(i)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for _, i := range []int{0, 1} {
			channel2 <- "Data ke-" + strconv.Itoa(i)
			time.Sleep(2 * time.Second)
		}
	}()

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1:", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data dari channel 2:", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}

		if counter == 2 {
			break
		}
	}

	fmt.Println("Selesai")
}
