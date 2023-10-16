package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	wg.Add(2)

	go buy(ch, &wg)
	go sell(ch, &wg)

	wg.Wait()

}

func sell(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("sending data")
	ch <- "hello welcome "
	fmt.Println("Data sended")
	wg.Done()
}
func buy(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for a data")
	val := <-ch
	fmt.Println("len:", len(val))
	fmt.Println("received ", val)
	wg.Done()
}
