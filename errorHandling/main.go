package main

import "fmt"

func main() {
	defer func(){
		if r:=recover();r!=nil{
			fmt.Println(r)
			fmt.Println("Recovered the situation")
		}
	}()
	fmt.Println("before defer")

MakeError()
fmt.Println("After defer")
}
func MakeError() {
	fmt.Println("next step is panic")
	panic("PANIC")
	fmt.Println("hello world")
}

func div(){
	
	a:=0
	res:=10/a
	fmt.Println(res)
}
