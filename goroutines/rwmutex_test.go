package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) Deposit(amount int) {
	account.RWMutex.Lock()
	defer account.RWMutex.Unlock()

	account.Balance += amount
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	defer account.RWMutex.RUnlock()

	return account.Balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.Deposit(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final balance:", account.GetBalance())
	// Output: Final balance: 10000
}
