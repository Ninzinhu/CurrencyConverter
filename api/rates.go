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

// Configura CORS para todas as rotas
func enableCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		
		// Headers CORS essenciais
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Responde imediatamente para requisições OPTIONS (pré-flight)
		if r.Method == "OPTIONS" {
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func getRates(w http.ResponseWriter, r *http.Request) {
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

	// Frontend estático (HTML/CSS/JS)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	// Aplica CORS a todas as rotas
	handler := enableCORS(r)

	// Inicia o servidor
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}