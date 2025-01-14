package quote

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DolarQuote struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type QuoteRepository interface {
	SaveQuote(quote *DolarQuote) error
}

func GetQuote(ctx context.Context) (*DolarQuote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result map[string]DolarQuote
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	quote, exists := result["USDBRL"]
	if !exists {
		log.Printf("Key 'USDBRL' not found in API response: %s", body)
		return nil, fmt.Errorf("failed to extract 'USDBRL' key from response: %w", err)
	}

	return &quote, nil
}

func SaveQuoteToDB(repo QuoteRepository) func(quote *DolarQuote) error {
	return func(quote *DolarQuote) error {
		return repo.SaveQuote(quote)
	}
}
