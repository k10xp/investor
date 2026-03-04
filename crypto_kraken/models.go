package crypto_kraken

type TickerResponse struct {
	Error  []string                `json:"error"`
	Result map[string]TickerRecord `json:"result"`
}

// https://docs.kraken.com/api/docs/rest-api/get-ticker-information/
type TickerRecord struct {
	Ask       []string `json:"a"` // ask [price, whole lot volume, lot volume]
	Bid       []string `json:"b"` // bid [price, whole lot volume, lot volume]
	LastTrade []string `json:"c"` // last trade closed [price, volume]
	Volume    []string `json:"v"` // volume [today, last 24 hours]
	VWAP      []string `json:"p"` // volume weighted avg price [today, last 24 hours]
	Trades    []int    `json:"t"` // number of trades [today, last 24 hours]
	Low       []string `json:"l"` // low [today, last 24 hours]
	High      []string `json:"h"` // high [today, last 24 hours]
	Open      string   `json:"o"` // today’s opening price
}
