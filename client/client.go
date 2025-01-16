package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Bid struct {
}

func main() {
	ctx := context.Background()
	bid, err := getBid(ctx)
	if err != nil {
		log.Fatalf("Error getting bid: %v", err)
	}
	err = saveBidToFile(bid, "cotacao.txt")
	if err != nil {
		log.Fatalf("Error saving bid to file: %v", err)
	}
	log.Println("Bid saved to bid.txt")

}

func saveBidToFile(bid, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Dólar: %s", bid)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func getBid(ctx context.Context) (string, error) {
	reqCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if reqCtx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("request timeout: %w", reqCtx.Err())
		}
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	log.Printf("API Response: %s", body)
	resp.Body = io.NopCloser(bytes.NewBuffer(body))

	decoder := json.NewDecoder(resp.Body)
	var bid string
	for {
		t, err := decoder.Token()
		if err != nil {
			break
		}
		if t == "bid" {
			decoder.Decode(&bid)
			break
		}
	}

	if bid == "" {
		return "", fmt.Errorf("failed to extract 'bid' key from response")
	}

	return bid, nil

}
