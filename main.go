package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "io/ioutil"
    "log"
)

type RateResponse struct {
    Rates map[string]float64 `json:"rates"`
    Base  string             `json:"base"`
}

func getRates(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    baseCurrency := vars["baseCurrency"]
    if baseCurrency == "" {
        baseCurrency = "USD" // Padrão para USD se não especificado
    }
    url := "https://api.exchangerate-api.com/v4/latest/" + baseCurrency

    response, err := http.Get(url)
    if err != nil {
        http.Error(w, "Erro ao obter taxas de câmbio", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        http.Error(w, "Erro ao ler resposta da API", http.StatusInternalServerError)
        return
    }
    var rates RateResponse
    json.Unmarshal(body, &rates)
    json.NewEncoder(w).Encode(rates)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/rates/{baseCurrency}", getRates).Methods("GET")

    http.Handle("/", r)
    log.Println("Servidor rodando na porta 8080")
    http.ListenAndServe(":8080", nil)
}
