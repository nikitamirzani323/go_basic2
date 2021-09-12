package main

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) unLock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.unLock()
	user2.unLock()
}

func main() {
	user1 := UserBalance{
		Name:    "Eko",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name:    "Edo",
		Balance: 1000000,
	}
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	time.Sleep(10 * time.Second)

	fmt.Println("User 1 :", user1.Name, ", Balance :", user1.Balance)
	fmt.Println("User 2 :", user2.Name, ", Balance :", user2.Balance)

	// go Transfer(&user2, &user1, 1000)

}
