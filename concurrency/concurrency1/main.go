package main

import (
	"fmt"
	"runtime"
	"time"
)

func calculateSqaure(i int, count *int) {
	fmt.Println(i * i)
	*count++

}
func main() {
	count := 0
	start := time.Now()
	for i := 1; i < 20000; i++ {
		//this means ten thousand function call exected in ten
		//ten thousand go routines.
		go calculateSqaure(i, &count)
	}
	defer fmt.Println(runtime.NumCPU(), ": Logical cpu used \n ")
	defer fmt.Println(runtime.NumGoroutine(), "Go rotines created \n ")
	time.Sleep(4 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("function took", elapsed )
	fmt.Println("")
}
