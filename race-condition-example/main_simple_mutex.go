package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// go run -race .
// func updateMessage(s string) {
// 	defer wg.Done()
// 	msg = s
// }

// use mutex
func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main_simple_mutex() {
	msg = "Hello, world!"

	// go run -race .
	// wg.Add(2)
	// go updateMessage("Hello, universe!")
	// go updateMessage("Hello, cosmos!")
	// wg.Wait()

	// use mutex
	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
