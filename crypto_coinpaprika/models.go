package crypto_coinpaprika

type Quote struct {
	Price float64 `json:"price"`
}

type Quotes struct {
	USD Quote `json:"USD"`
}

type Ticker struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Rank   int    `json:"rank"`
	Quotes Quotes `json:"quotes"`
}
