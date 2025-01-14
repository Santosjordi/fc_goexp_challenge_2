package main

import (
	"log"
	"net/http"

	"github.com/santosjordi/posgoexp/challenges/ctx-client-server/handler"
)

func main() {

	http.HandleFunc("/cotacao", handler.QuoteHandler)
	log.Println("Server started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
