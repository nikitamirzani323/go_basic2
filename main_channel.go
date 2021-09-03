package main

import "fmt"

func main() {
	c := make(chan string)

	go func(n string) {
		c <- n
	}("Peler")

	var chanFunc = func(n string) {
		c <- n
	}

	go chanFunc("Anjing")

	b := <-c
	o := <-c
	fmt.Println(b)
	fmt.Println(o)

	chankedua := make(chan string)
	go CetakNama(chankedua, "hahahahhahaha")
	fmt.Println(<-chankedua)

	pp := make(chan int, 2)
	go func() {
		for p := range pp {
			p = <-pp
			fmt.Println("Terima Data ======", p)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Kirim data")
		pp <- i
	}
}

func CetakNama(c chan string, v string) {
	c <- v
}
