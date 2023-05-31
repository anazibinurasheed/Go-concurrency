package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mut sync.Mutex
	abc := 10
	go one(&abc,&mut)
	go two(&abc,&mut)
	go three(&abc,&mut)
	time.Sleep(2 * time.Second)
	fmt.Println(abc)
}

func one(abc *int,m *sync.Mutex) {
m.Lock()
	*abc = 00
	m.Unlock()
}
func two(abc *int,m *sync.Mutex) {
	m.Lock()
	*abc = 7888
	m.Unlock()
}
func three(abc *int,m *sync.Mutex) {
	m.Lock()
	*abc = 8000
m.Unlock()
}
