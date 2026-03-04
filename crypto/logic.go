package crypto

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

func FetchTickers(url string) ([]Ticker, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tickers []Ticker
	if err := json.Unmarshal(body, &tickers); err != nil {
		return nil, err
	}

	return tickers, nil
}

func ExportCSV(filename string, tickers []Ticker) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// csv header
	if err := w.Write([]string{"rank", "name", "symbol", "price_usd"}); err != nil {
		return err
	}

	for _, t := range tickers {
		record := []string{
			strconv.Itoa(t.Rank),
			t.Name,
			t.Symbol,
			strconv.FormatFloat(t.Quotes.USD.Price, 'f', 8, 64),
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return w.Error()
}
