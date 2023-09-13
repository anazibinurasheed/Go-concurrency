package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	b := 0
	a := 10 / b
	fmt.Println(a)

}

