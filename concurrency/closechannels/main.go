package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 88
	ch <- 99
	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)
	val,ok=<-ch
	fmt.Println(val,ok)
	val,ok=<-ch
	fmt.Println(val,ok)
}
