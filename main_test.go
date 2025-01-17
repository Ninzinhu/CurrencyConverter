package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "encoding/json"
)

func TestGetRates(t *testing.T) {
    req, err := http.NewRequest("GET", "/api/rates", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(getRates)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var response RateResponse
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Errorf("handler returned invalid JSON: %v", err)
    }

    if response.Base != "JPY" {
        t.Errorf("handler returned unexpected base currency: got %v want %v", response.Base, "JPY")
    }

    if len(response.Rates) == 0 {
        t.Errorf("handler returned empty rates map")
    }
}
