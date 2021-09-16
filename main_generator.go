package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var mutex sync.RWMutex

func generatornomor(wId int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {

	for capture := range jobs {
		log.Printf("=======Worker ID: %d ===========\n", wId)

		// retry jika ada error
		for {
			var outerError error
			func(outerError *error) {

				results <- capture

			}(&outerError)
			if outerError == nil {
				break
			}
		}

	}
	wg.Done()
}
func main() {
	fmt.Println("Worker pools in go")
	runtime.GOMAXPROCS(10)
	var count_jobs int = 0
	var count_receive int = 0
	start := time.Now()
	totals := 9999
	worker := 10
	buffer := totals + 1
	jobs := make(chan string, buffer)
	results := make(chan string, buffer)
	wg := &sync.WaitGroup{}

	for w := 0; w < worker; w++ {
		wg.Add(1)
		mutex.Lock()
		go generatornomor(w, jobs, results, wg)
		mutex.Unlock()
	}
	for i := 0; i <= totals; i++ {
		count_jobs++
		mutex.Lock()
		if i >= 0 && i < 10 {
			jobs <- "000" + strconv.Itoa(i)
		}
		if i >= 10 && i <= 99 {
			jobs <- "00" + strconv.Itoa(i)
		}
		if i >= 100 && i <= 999 {
			jobs <- "0" + strconv.Itoa(i)
		}
		if i > 999 && i <= 9999 {
			jobs <- strconv.Itoa(i)
		}
		mutex.Unlock()
	}
	close(jobs)
	wg.Wait()
	for i := 0; i <= totals; i++ {
		count_receive++
		log.Println(<-results)
	}
	time.Sleep(2 * time.Millisecond)
	log.Println("Kirim Jobs : ", count_jobs)
	log.Println("Terima Jobs : ", count_receive)
	log.Println("time : ", time.Since(start))
}
