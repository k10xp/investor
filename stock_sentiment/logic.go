package stock_sentiment

import (
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func FetchSentiment() ([]RedditStock, error) {
	url := "https://api.tradestie.com/v1/apps/reddit" //?date default to today

	//certificate expired, temp workaround to skip TLS verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var stocks []RedditStock
	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	return stocks, nil
}

func ExportCSV(path string, stocks []RedditStock) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	//header
	if err := w.Write([]string{
		"ticker",
		"no_of_comments",
		"sentiment",
		"sentiment_score",
	}); err != nil {
		return fmt.Errorf("write header failed: %w", err)
	}

	//rows
	for _, s := range stocks {
		row := []string{
			s.Ticker,
			strconv.Itoa(s.NoOfComments),
			s.Sentiment,
			strconv.FormatFloat(s.SentimentScore, 'f', 6, 64),
		}
		if err := w.Write(row); err != nil {
			return fmt.Errorf("write row failed: %w", err)
		}
	}

	if err := w.Error(); err != nil {
		return fmt.Errorf("csv writer error: %w", err)
	}

	return nil
}
