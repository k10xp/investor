# Finance and Crypto investor

Fetch crypto ticker data and export as csv file. Library and standalone script.

Run

```shell
go run main.go
```

## Example

```go
func main() {
	//fetch
	coinpaprika_tickers, c_fetch_err := crypto_coinpaprika.FetchTickers()
	if c_fetch_err != nil {
		log.Fatalf("error fetching ticker data: %s", c_fetch_err)
	}

	//export to csv
	c_save_err := crypto_coinpaprika.ExportCSV("./data/coinpaprika_export.csv", coinpaprika_tickers)
	if c_save_err != nil {
		log.Fatalf("export csv error: %s", c_save_err)
	}
}
```

```go
func main(){
    //fetch
	kraken_tickers, err := crypto_kraken.FetchTickers()
	if err != nil {
		log.Fatalf("error fetching ticker data: %s", err)
	}

	//export to csv
	k_save_err := crypto_kraken.ExportCSV("./data/kraken_export.csv", kraken_tickers)
	if k_save_err != nil {
		log.Fatalf("export csv error: %s", k_save_err)
	}
}
```
