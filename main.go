package main

import "eshaanagg/go/cryptomasters/api"

func main() {
	rate, err := api.GetRate("BTC")
	print(rate, err)	
}