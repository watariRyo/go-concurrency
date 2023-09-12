package main

import (
	"fmt"
	"sync"
)

// func printSomething(s string) {
// 	fmt.Println(s)
// }

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

// goroutines
func main() {
	// wait a moment. bad way.
	// go printSomething("This is the 1st thing to be printed!")
	// time.Sleep(1 * time.Second)
	// printSomething("This is the 2nd thing to be printed!")
	// printSomething("This is the 3rd thing to be printed!")

	// waitGroups
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epcilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()
}
