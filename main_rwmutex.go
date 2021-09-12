package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) getBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func main() { //TUJUAN UNTUK SUPAYA URUT DAN TIDAK DEADLOCK
	account := BankAccount{}

	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				account.AddBalance(1)
				fmt.Println(account.getBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total balance : ", account.getBalance())
}
