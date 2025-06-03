package main

import (
	"fmt"
	"sync"

	"github.com/kaipov24/cryptomsters/api"
)

func main() {
	currencies := []string{"BTC", "ETH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The current price of %v is %.8f \n", rate.Currency, rate.Price)
	}
}
