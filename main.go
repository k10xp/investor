package main

import (
	"fmt"
	"log"

	"github.com/k10xp/investor/crypto"
)

func main() {
	url := "https://api.coinpaprika.com/v1/tickers?quotes=USD"
	tickers, err := crypto.FetchTickers(url)
	if err != nil {
		log.Fatal("error fetching ticker data")
	}

	//top 5 entries
	for i, t := range tickers {
		if i >= 5 {
			break
		}
		fmt.Printf("%d. %s (%s): $%.2f\n", t.Rank, t.Name, t.Symbol, t.Quotes.USD.Price)
	}
}
