package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	usecases "github.com/3Thiago/Client-Server-API/Server/Usecases"
)

func GetPricepHandler(w http.ResponseWriter, r *http.Request) {
	currencyData, error := usecases.GetCurrencyData()
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprint(currencyData.USDBRL.Bid))
}
