package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		//the go routine will take the variable value when the main
		//routine wait for it
		//if we want deterministic answer give the value as function argument .

		//else it will also show a race condition
		go func() {
			fmt.Println(i)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("All done")
}
