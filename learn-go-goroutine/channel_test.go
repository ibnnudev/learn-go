package learngogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ibnu Sutio"
		fmt.Println("Selesai mengirim data")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
