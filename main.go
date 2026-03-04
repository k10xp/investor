package main

import (
	"log"
	"sync"

	"github.com/k10xp/investor/crypto_coinpaprika"
	"github.com/k10xp/investor/crypto_kraken"
)

func main() {
	//setup go routine
	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	//CoinPaprika
	wg.Add(1)
	go func() {
		defer wg.Done()

		tickers, err := crypto_coinpaprika.FetchTickers()
		if err != nil {
			errChan <- err
			return
		}

		crypto_coinpaprika.TopEntries(tickers, 5)

		if err := crypto_coinpaprika.ExportCSV("./data/coinpaprika_export.csv", tickers); err != nil {
			errChan <- err
		}
	}()

	//Kraken
	wg.Add(1)
	go func() {
		defer wg.Done()

		tickers, err := crypto_kraken.FetchTickers()
		if err != nil {
			errChan <- err
			return
		}

		if err := crypto_kraken.ExportCSV("./data/kraken_export.csv", tickers); err != nil {
			errChan <- err
		}
	}()

	//wait for go routine finish + handle all errors
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			log.Fatalf("Error occurred: %s", err)
		}
	}

	log.Println("All tasks completed successfully.")
}
