package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 3)
	defer close(channel)

	go In(channel)
	go Out(channel)

	time.Sleep(3 * time.Second)
}

func In(channel chan<- string) {
	channel <- "Eko Kurniawan"
	channel <- "Shinta"
	channel <- "Bernad"
	fmt.Println("Kirim Data Selesai")
}
func Out(channel <-chan string) {
	data1 := <-channel
	data2 := <-channel
	data3 := <-channel
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)
	fmt.Println("Terima Data Selesai")
}
