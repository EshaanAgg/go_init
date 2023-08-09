package main

import (
	"eshaanagg/go/cryptomasters/api"
	"fmt"
)

func main() {
	rate, err := api.GetRate("BTC")

	if err == nil {
		fmt.Printf("The rate for the %v is %.2f\n", rate.Currency, rate.Price)
	}
}