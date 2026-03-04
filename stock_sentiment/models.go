package stock_sentiment

type RedditStock struct {
	NoOfComments   int     `json:"no_of_comments"`
	Sentiment      string  `json:"sentiment"`
	SentimentScore float64 `json:"sentiment_score"`
	Ticker         string  `json:"ticker"`
}
