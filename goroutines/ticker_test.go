package main

import (
	"fmt"
	"testing"
	"time"
)

// test ticker
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		case <-done:
			fmt.Println("Ticker stopped")
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			return
		}
	}
}

func TestTick(t *testing.T) {
	tick := time.Tick(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-tick:
			fmt.Println(t)
		case <-done:
			fmt.Println("Tick stopped")
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			return
		}
	}
}

func TestTickerStop(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		case <-done:
			fmt.Println("Ticker stopped")
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
			return
		}
	}
}
