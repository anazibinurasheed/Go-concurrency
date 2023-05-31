package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go fOne(ch1)
	go fTwo(ch2)
	time.Sleep(1*time.Second)

	select {
	case val := <-ch1:
		fmt.Println(val)

	case val := <-ch2:
		fmt.Println(val)
	// case time.After(1*time.Second): 
	// fmt.Println("hello world ")

	// default:
	// 	fmt.Println("I am default statement")
	}
}

func fTwo(ch2 chan string) {

	ch2 <- "channel two "
}

func fOne(ch1 chan string) {
	ch1 <- "channel one"
}
