package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	usecases "github.com/3Thiago/Client-Server-API/Server/Usecases"
)

func GetPricepHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	currencyData, error := usecases.GetCurrencyData(ctx)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprint(currencyData.USDBRL.Bid))
}
