package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Site struct {
	URL string
}
type Result struct {
	Status      int
	Linkwebsite string
}

func crawl(wId int, jobs <-chan Site, results chan<- Result, wg *sync.WaitGroup) {

	for site := range jobs {
		log.Printf("=======Worker ID: %d ===========\n", wId)
		resp, err := http.Get(site.URL)

		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode, Linkwebsite: site.URL}
	}
	wg.Done()
}
func main() {
	fmt.Println("Worker pools in go")
	worker := 10
	totals := 88
	start := time.Now()
	jobs := make(chan Site, 90)
	results := make(chan Result, 90)
	wg := &sync.WaitGroup{}

	for w := 0; w < worker; w++ {
		wg.Add(1)
		go crawl(w, jobs, results, wg)
	}

	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://detik.com",
		"https://kompas.com",
		"https://youtube.com",
		"https://stackoverflow.com/",
		"https://morioh.com/",
		"https://dasarpemrogramangolang.novalagung.com/",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)
	wg.Wait()

	for a := 0; a < totals; a++ {
		result := <-results
		log.Println(result)
	}

	elsapsed := time.Since(start)
	fmt.Println("Waktu :", elsapsed)
}
