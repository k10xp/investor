package crypto

import (
	"encoding/json"
	"io"
	"net/http"
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
