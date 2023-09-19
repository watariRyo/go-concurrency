package main

import (
	"fmt"
	"sync"
	"time"
)

// This is a simple implementation of Dijkstra's solution to the "Dining Philosophers" dilemma.
// There are five philosophers and five forks.

// Philosopher is a struct which stores information about a philosoper
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of all philosopers
var philosopers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotele", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define
var hunger = 3 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("--------------------------------")
	fmt.Println("The table is empty.")

	// start a meal
	dine()

	fmt.Println("The table is empty.")
}

func dine() {
	// eatTime = 0 * time.Second
	// sleepTime = 0 * time.Second
	// thinkTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philosopers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosopers))

	// forks is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosopers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosopers); i++ {
		// fire off a goroutine for the current philosoper
		go diningProblem(philosopers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosoper Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosoper.name)
	seated.Done()

	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		if philosoper.leftFork > philosoper.rightFork {
			forks[philosoper.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosoper.name)
			forks[philosoper.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosoper.name)
		} else {
			forks[philosoper.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosoper.name)
			forks[philosoper.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosoper.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosoper.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosoper.name)
		time.Sleep(thinkTime)

		forks[philosoper.leftFork].Unlock()
		forks[philosoper.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosoper.name)
	}

	fmt.Println(philosoper.name, " is satisfied.")
	fmt.Println(philosoper.name, " left the table.")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosoper.name)
	orderMutex.Unlock()
}
