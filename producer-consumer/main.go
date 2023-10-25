package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch

}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved Order number %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed += 1

		} else {
			pizzasMade += 1
		}
		total += 1
		fmt.Printf("Making pizza number %d, it will take %d seconds\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay * time.Now().Second()))

		if rnd <= 2 {
			msg = fmt.Sprintf("*** WE RAN OUT OF INGREDIENTS FOR PIZZA NUMBER %d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** COOKS ON STRIKE, COULDN'T MAKE PIZZA NUMBER %d", pizzaNumber)

		} else {
			success = true
			msg = fmt.Sprintf("Pizza order %d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}

}

func pizzaria(pizzaMaker *Producer) {
	// keep track of which pizza we are trying to make
	var i = 0

	// run forever or until we recieve a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		// try to make a pizza

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried making a pizza (we sent data to the data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// need to close chanel
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}

		// decision
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out message
	color.Cyan("The Pizzaria is open for Buisness")
	color.Cyan("---------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzaria(pizzaJob)
	// create and run a consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order number %d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Customer is Mad :(")
			}
		} else {
			color.Cyan("Done Making Pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing Channel, %s", err)
			}
		}
	}

	// print out the ending message

}
