package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Package main implements a simulation to illustrate the Dining Philosophers problem.

Problem Statement:
-----------------

Five philosophers sit around a circular dining table. Each philosopher thinks deeply and occasionally interrupts their thinking to eat from a bowl of spaghetti placed in front of them. A fork is placed between each pair of adjacent philosophers, and each philosopher may only use the forks on their immediate left and right.

The philosophers must pick up both forks before eating and put them down after eating. A philosopher can only pick up one fork at a time, and they must pick up and put down forks in such a way that deadlock or contention does not occur.

Constraints:
------------
1. A philosopher can either think or eat, but not both at the same time.
2. A philosopher must pick up both left and right forks to eat.
3. A philosopher must put down both forks after eating.
4. A fork can be held by only one philosopher at a time.

Objective:
----------
The objective is to come up with a protocol that allows the philosophers to share the forks as efficiently as possible, without running into problems such as deadlocks, livelocks, or contention.

Author: Your Name
Date: YYYY-MM-DD
*/

// philosopher struct that stores informatino about a philosopher
type Philosopher struct {
	name   string
	r_fork int
	l_fork int
}

// philosophers list of philosophers
var philosophers = []Philosopher{
	{name: "Plato", l_fork: 4, r_fork: 0},
	{name: "Socrates", l_fork: 0, r_fork: 1},
	{name: "Aristotle", l_fork: 1, r_fork: 2},
	{name: "Pascal", l_fork: 2, r_fork: 3},
	{name: "Locke", l_fork: 3, r_fork: 4},
}

// define a few more varaibles
var hunger = 3 // how many times does a person eat?
var eatTime = 1 * time.Second
var thinkTime = 1 * time.Second
var sleepTime = 1 * time.Second

// define left philosophers
var leaving []string

func main() {
	// Print a welcome message
	fmt.Println("Dining Philosophers problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")

	// print out the order in which philosophers leave
	fmt.Println("the final order in which the philospher's left:")
	fmt.Println("-----------------------------------------------")
	for index, value := range leaving {
		fmt.Printf("\t%d. %s\n", index+1, value)
	}
}

func dine() {

	// eatTime = 0 * time.Second
	// sleepTime = 0 * time.Second
	// thinkTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var leaving_phil sync.Mutex

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a go routine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated, &leaving_phil)
	}

	wg.Wait()

}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup, leaving_phil *sync.Mutex) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {

		// get lock on both forks
		if philosopher.l_fork > philosopher.r_fork {
			forks[philosopher.r_fork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.l_fork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.l_fork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.r_fork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.l_fork].Unlock()
		forks[philosopher.r_fork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, " is satisfied.")
	leaving_phil.Lock()
	leaving = append(leaving, philosopher.name)
	leaving_phil.Unlock()
	fmt.Println(philosopher.name, " left the table.")

}
