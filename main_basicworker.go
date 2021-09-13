package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Worker Pool")
	start := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// ini running mau berapa worker
	for i := 1; i <= 500; i++ {
		go consumer(i, jobs, results)
	}

	// kirim job ke producer
	go producer(jobs, 10000)

	// ambil result
	for i := 1; i <= 10000; i++ {
		res := <-results
		fmt.Println("Hasil ke ", res)
	}

	fmt.Println("==============")
	elsapsed := time.Since(start)
	fmt.Println("Waktu :", elsapsed)

	totalCPU := runtime.NumCPU()
	fmt.Println("CPU", totalCPU)

	// -1 artinya dibawah 0 THREAD SESUAI PC
	// 1 artinya diatas 0 Mengubah Jumlah Thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)
}

func FakeHttpRequest(x int) int {
	fmt.Printf("INSERT INTO %d\n", x)
	return x
}

func producer(jobs chan<- int, size int) {
	for i := 1; i <= size; i++ {
		jobs <- i
	}
	close(jobs)
}
func consumer(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Consumer ke ", id, " Mulai")

		time.Sleep(time.Millisecond * 100)

		results <- FakeHttpRequest(job)
	}
}
