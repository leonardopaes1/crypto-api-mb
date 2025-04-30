package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Estruturas para receber o JSON da API externa
type TickerData struct {
	Ticker struct {
		High  string `json:"high"`
		Low   string `json:"low"`
		Vol   string `json:"vol"`
		Last  string `json:"last"`
		Buy   string `json:"buy"`
		Sell  string `json:"sell"`
		Open  string `json:"open"`
		Date  int64  `json:"date"`
		Pair2 string `json:"pair"`
	} `json:"ticker"`
}

type Response struct {
	Environment string        `json:"environment"`
	Ticker      []TickerEntry `json:"ticker"`
}

type TickerEntry struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
	Open string `json:"open"`
	Date int64  `json:"date"`
	Pair string `json:"pair"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "dev"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pair := r.URL.Path[1:]

		if pair == "" {
			http.Error(w, "Missing cryptocurrency pair. Please use /{pair}, e.g., /BTC", http.StatusBadRequest)
			return
		}

		url := fmt.Sprintf("https://www.mercadobitcoin.net/api/%s/ticker/", pair)
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Error fetching ticker: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Novo tratamento: se a resposta da API externa não for 200, já retorna erro
		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("Failed to fetch ticker for pair %s. HTTP status: %d", pair, resp.StatusCode), http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var tickerData TickerData
		if err := json.Unmarshal(body, &tickerData); err != nil {
			http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}

		response := Response{
			Environment: environment,
			Ticker: []TickerEntry{
				{
					High: tickerData.Ticker.High,
					Low:  tickerData.Ticker.Low,
					Vol:  tickerData.Ticker.Vol,
					Last: tickerData.Ticker.Last,
					Buy:  tickerData.Ticker.Buy,
					Sell: tickerData.Ticker.Sell,
					Open: tickerData.Ticker.Open,
					Date: tickerData.Ticker.Date,
					Pair: tickerData.Ticker.Pair2,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Adicionando o endpoint de healthcheck
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
