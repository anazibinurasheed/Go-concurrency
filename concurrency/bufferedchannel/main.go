package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	wg.Add(2)

	go sell(ch, &wg)

	wg.Wait()

}
func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	ch <- 3
	go buy(ch, wg)

	fmt.Println("", ch)
	fmt.Println("sent all data to the channel")
	wg.Done()
}

func buy(ch chan int, wg *sync.WaitGroup) {

	fmt.Println("Waiting for data")
	fmt.Println("len:", len(ch), "cap:", cap(ch))

	fmt.Println("recieved data from the channel", <-ch)
	fmt.Println("recieved data from the channel", <-ch)
	fmt.Println("recieved data from the channel", <-ch)
	wg.Done()

}
