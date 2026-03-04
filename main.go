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
		log.Fatalf("error fetching ticker data: %s", err)
	}

	//top 5 entries
	for i, t := range tickers {
		if i >= 5 {
			break
		}
		_, _ = fmt.Printf("%d. %s (%s): $%.2f\n", t.Rank, t.Name, t.Symbol, t.Quotes.USD.Price)
	}

	//export to csv
	err1 := crypto.ExportCSV("./data/coinpaprika_export.csv", tickers)
	if err1 != nil {
		log.Fatalf("export csv error: %s", err1)
	}
}
