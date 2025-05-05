package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Name    string
	Balance int
	Mutex   sync.Mutex
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	if fmt.Sprintf("%p", user1) < fmt.Sprintf("%p", user2) {
		user1.Lock()
		defer user1.Unlock()
		fmt.Println("Lock User 1:", user1.Name)

		time.Sleep(1 * time.Second)

		user2.Lock()
		defer user2.Unlock()
		fmt.Println("Lock User 2:", user2.Name)

		user1.Change(-amount)
		user2.Change(amount)
	} else {
		user2.Lock()
		defer user2.Unlock()
		fmt.Println("Lock User 2:", user2.Name)

		time.Sleep(1 * time.Second)

		user1.Lock()
		defer user1.Unlock()
		fmt.Println("Lock User 1:", user1.Name)

		user2.Change(amount)
		user1.Change(-amount)
	}

	time.Sleep(1 * time.Second)
}

func TestDeadlockPrevention(t *testing.T) {
	user1 := UserBalance{
		Name:    "Ibnu",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Nada",
		Balance: 1000000,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go Transfer(&user1, &user2, 100000, &wg)
	go Transfer(&user2, &user1, 200000, &wg)

	wg.Wait()

	fmt.Println("User ", user1.Name, " | Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, " | Balance ", user2.Balance)
}
