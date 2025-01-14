package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/santosjordi/posgoexp/challenges/ctx-client-server/quote"
)

func QuoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	log.Println("Request started")
	defer log.Println("Request ended")

	quote, err := quote.GetQuote(ctx)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			w.WriteHeader(http.StatusGatewayTimeout)
			w.Write([]byte("504 Gateway Timeout"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	response := map[string]string{"bid": quote.Bid}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
