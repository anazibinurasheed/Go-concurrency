package main

import (
	"fmt"
	"time"
)
func main() {
	go func() {
		fmt.Println("in anonymous function")
	}()
time.Sleep(1* time.Second)
}
