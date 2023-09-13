package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func calculateSqaure(i int, count *int, mutex *sync.Mutex) {
	fmt.Println(i * i)
	mutex.Lock()
	*count++
	mutex.Unlock()

}
func main() {
	count := 0
	start := time.Now()
	mutex := sync.Mutex{}

	for i := 1; i < 200; i++ {
		//this means n function call  will execute in
		//n goroutines.
		go calculateSqaure(i, &count, &mutex)
	}
	fmt.Println(runtime.NumCPU(), ": Logical cpu used \n ")
	fmt.Println(runtime.NumGoroutine(), "Goroutines created \n ")

	time.Sleep(4 * time.Second)

	elapsed := time.Since(start)

	fmt.Println("function took", elapsed)
}
