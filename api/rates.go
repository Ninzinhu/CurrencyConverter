package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type RateResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

// Handler principal para a Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Configura CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	// Log para debug
	println("Method:", r.Method)
	println("URL:", r.URL.String())

	// Responde para requisições OPTIONS (pré-flight)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Extrai a moeda base da URL
	// URL será: /api/rates?base=USD ou /api/rates/USD
	baseCurrency := r.URL.Query().Get("base")
	if baseCurrency == "" {
		// Tenta extrair da path
		path := strings.TrimPrefix(r.URL.Path, "/api/rates/")
		if path != "" && path != "/api/rates" {
			baseCurrency = path
		} else {
			baseCurrency = "BRL" // Default
		}
	}

	// Chama a API externa
	url := "https://api.exchangerate-api.com/v4/latest/" + baseCurrency
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erro ao acessar API de câmbio", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Erro ao ler resposta", http.StatusInternalServerError)
		return
	}

	var rates RateResponse
	if err := json.Unmarshal(body, &rates); err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rates)
}