package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go CetakNama("Hi")
	go CetakNama("Hello")

	go CetakNama("Whats up")
	wg.Wait()
}

func CetakNama(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
