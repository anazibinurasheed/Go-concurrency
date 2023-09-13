package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func calculateSqaure(i int, count *int) {
	fmt.Println(i * i)
	mutex := sync.Mutex{}
	mutex.Lock()
	*count++
	mutex.Unlock()

}
func main() {
	count := 0
	start := time.Now()
	for i := 1; i < 200; i++ {
		//this means n function call  will execute in
		//n goroutines.
		go calculateSqaure(i, &count)
	}
	defer fmt.Println(runtime.NumCPU(), ": Logical cpu used \n ")
	defer fmt.Println(runtime.NumGoroutine(), "Goroutines created \n ")
	time.Sleep(4 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("function took", elapsed)
	fmt.Println("")
}
