package main

import "fmt"

func main() {
	defer func(){
		r:=recover()
		if r!= nil {
			fmt.Println(r)

		}
	}()
test(78)

}

func test(i interface{}){
	switch v:=i.(type){
	case int : fmt.Printf("%T %#v ",v,v)
	}
}
