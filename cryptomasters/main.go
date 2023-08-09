package main

import (
	"eshaanagg/go/cryptomasters/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "BCH", "ETH"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currency string) {
			getCurrencyData(currency)
			wg.Done()
		}(currency)
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f USD.\n", rate.Currency, rate.Price)
	}
}
