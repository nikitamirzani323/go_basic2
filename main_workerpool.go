package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}
type Result struct {
	Status      int
	Linkwebsite string
}

func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)

		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode, Linkwebsite: site.URL}
	}
}
func main() {
	fmt.Println("Worker pools in go")

	jobs := make(chan Site, 22)
	results := make(chan Result, 22)

	for w := 1; w <= 22; w++ {
		go crawl(w, jobs, results)
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
		"https://www.viva.co.id/",
		"https://www.okezone.com/",
		"https://tavuli.com/?token=qC5Ym%2BBvXzabGp34jJlKvnC6wCrr3pLCwBzsLoSzl4k%3D",
		"https://webtor.io/",
		"https://santekno.com/",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= 22; a++ {
		result := <-results
		log.Println(result)
	}
}
