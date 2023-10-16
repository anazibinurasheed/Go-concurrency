package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go sendValue(ch)
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-time.After(1 * time.Second):
		fmt.Println("TimeOut")

	}
	//the select statement will wait for 1 second for receive
	//operation on channel after 1 sec time.After statement will
	//execute.
}

func sendValue(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 88

}
