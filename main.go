package main

import (
	"fmt"

	"github.com/kaipov24/cryptomsters/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("The current price of %v is %.8f \n", rate.Currency, rate.Price)
	}
	fmt.Println(rate.Price, err)
}
