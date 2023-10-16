package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	wg.Add(1)
	go buy(ch, &wg)

	ch <- 100
	ch <- 200
	ch <- 300
	close(ch)
	wg.Wait()

}

func buy(ch chan int, wg *sync.WaitGroup) {
	for val := range ch {
		fmt.Println(val)

	}
	wg.Done()
}
