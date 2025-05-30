package main

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type RateResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getRates(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	baseCurrency := vars["baseCurrency"]
	if baseCurrency == "" {
		baseCurrency = "BRL"
	}

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

func main() {
	r := mux.NewRouter()
	
	// API
	r.HandleFunc("/api/rates/{baseCurrency}", getRates).Methods("GET", "OPTIONS")
	
	// Frontend estático (com fallback para SPA)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))
	
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}