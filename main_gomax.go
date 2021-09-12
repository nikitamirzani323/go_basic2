package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("CPU", totalCPU)

	// -1 artinya dibawah 0 THREAD SESUAI PC
	// 1 artinya diatas 0 Mengubah Jumlah Thread
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)

	group.Wait()
}
