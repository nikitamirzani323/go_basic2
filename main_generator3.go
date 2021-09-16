package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	var count_jobs int = 0
	totals := 9999
	start := time.Now()
	for i := 0; i <= totals; i++ {
		count_jobs++

		if i >= 0 && i < 10 {
			log.Printf("000%s\n", strconv.Itoa(i))
		}
		if i >= 10 && i <= 99 {
			log.Printf("00%s\n", strconv.Itoa(i))
		}
		if i >= 100 && i <= 999 {
			log.Printf("0%s\n", strconv.Itoa(i))
		}
		if i > 999 && i <= 9999 {
			log.Printf("%s\n", strconv.Itoa(i))
		}
	}
	log.Println("total : ", count_jobs)
	log.Println("Time : ", time.Since(start))
}
