package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int

	// mutex for lock balance
	var balance sync.Mutex

	// print starting values
	fmt.Printf("Initial Account Balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revanue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "PartTimeJob", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made ; keep a running total
	for i, v := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				balance.Unlock()

				fmt.Printf("On week %d we earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}

		}(i, v)
	}

	wg.Wait()

	// print final balance

	fmt.Printf("Final bank Balance: $%d.00", bankBalance)
	fmt.Println()

}
