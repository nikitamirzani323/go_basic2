package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}
	group.Wait()
	fmt.Println("Complete")
}
func RunAsync(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}
