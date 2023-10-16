package main

import "fmt"

// Send only chan and receive only chan
func main() {

	ch := make(chan interface{})
	go ReceiveOnlyChan(ch)
	go SendOnlyChan(ch)
	ch <- "a message from entry point"
	fmt.Println(<-ch)
}
func SendOnlyChan(ch chan<- interface{}) {

	ch <- "A message from SendOnlyChan"
}

// example of receive only chan
func ReceiveOnlyChan(ch <-chan interface{}) {

	fmt.Println(<-ch)

}
