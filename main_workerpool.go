package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
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

		// retry jika ada error
		for {
			var outerError error
			func(outerError *error) {
				defer func() {
					if err := recover(); err != nil {
						*outerError = fmt.Errorf("%v", err)
					}
				}()

				resp, err := http.Get(site.URL)

				if err != nil {
					log.Println(err.Error())
				}
				results <- Result{Status: resp.StatusCode, Linkwebsite: site.URL}
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
	totals := 88
	worker := totals
	buffer := totals + 1
	start := time.Now()
	jobs := make(chan Site, buffer)
	results := make(chan Result, buffer)
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
