package crypto_coinpaprika

import "fmt"

func TopEntries(tickers []Ticker, num_entries int) {
	for i, t := range tickers {
		if i >= num_entries {
			break
		}
		_, _ = fmt.Printf("%d. %s (%s): $%.2f\n", t.Rank, t.Name, t.Symbol, t.Quotes.USD.Price)
	}
}
