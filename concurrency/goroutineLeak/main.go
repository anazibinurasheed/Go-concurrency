package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go leak(&wg)

	wg.Wait()

}

func leak(wg *sync.WaitGroup) {

	ch := make(chan string, 3)

	go func() {
		v := <-ch
		fmt.Println("executed", v)
		wg.Done()
	}()
	
	fmt.Println("exiting from leak method")

	wg.Done()
}
