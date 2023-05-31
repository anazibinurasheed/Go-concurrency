package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main() {
	start:=time.Now()
	wg.Add(10)
	for i:=0;i<10;i++{
go calculatesquare(i,&wg)
	}
wg.Wait()
	fmt.Println("time took",time.Since(start))
	

}

func calculatesquare(i int,wg *sync.WaitGroup){
	fmt.Println(i*i)
	wg.Done()
}