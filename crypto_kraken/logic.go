package crypto_kraken

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func FetchTickers() (*TickerResponse, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.kraken.com/0/public/Ticker", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tr TickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return nil, err
	}
	if len(tr.Error) > 0 {
		return nil, fmt.Errorf("kraken error: %v", tr.Error)
	}

	return &tr, nil
}

func ExportCSV(filename string, tr *TickerResponse) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"pair",
		"ask_price", "ask_whole_lot", "ask_lot",
		"bid_price", "bid_whole_lot", "bid_lot",
		"last_price", "last_volume",
		"vol_today", "vol_24h",
		"vwap_today", "vwap_24h",
		"trades_today", "trades_24h",
		"low_today", "low_24h",
		"high_today", "high_24h",
		"open",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for pair, rec := range tr.Result {
		// guard against missing elements
		a := pad(rec.Ask, 3)
		b := pad(rec.Bid, 3)
		c := pad(rec.LastTrade, 2)
		v := pad(rec.Volume, 2)
		p := pad(rec.VWAP, 2)
		l := pad(rec.Low, 2)
		h := pad(rec.High, 2)

		t0, t1 := "", ""
		if len(rec.Trades) > 0 {
			t0 = strconv.Itoa(rec.Trades[0])
		}
		if len(rec.Trades) > 1 {
			t1 = strconv.Itoa(rec.Trades[1])
		}

		row := []string{
			pair,
			a[0], a[1], a[2],
			b[0], b[1], b[2],
			c[0], c[1],
			v[0], v[1],
			p[0], p[1],
			t0, t1,
			l[0], l[1],
			h[0], h[1],
			rec.Open,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}

	return w.Error()
}

func pad(s []string, n int) []string {
	out := make([]string, n)
	for i := 0; i < n; i++ {
		if i < len(s) {
			out[i] = s[i]
		} else {
			out[i] = ""
		}
	}

	return out
}
