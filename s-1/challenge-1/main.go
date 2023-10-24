package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"
	var wg sync.WaitGroup 

	wg.Add(3)

	go updateMessage("Hello, universe!", &wg)
	printMessage()

	go updateMessage("Hello, cosmos!", &wg)
	printMessage()

	go updateMessage("Hello, world!", &wg)
	printMessage()

	wg.Wait()
}
