package usecases

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	database "github.com/3Thiago/Client-Server-API/Server/Database"
	model "github.com/3Thiago/Client-Server-API/Server/Model"
)

func GetCurrencyData(ctx context.Context) (*model.CurrencyData, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	resp, error := http.DefaultClient.Do(req)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var currencyData model.CurrencyData
	error = json.Unmarshal(body, &currencyData)
	if error != nil {
		return nil, error
	}
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	InsertCurrencyData(ctx, &currencyData)
	return &currencyData, nil
}

func InsertCurrencyData(ctx context.Context, currencyData *model.CurrencyData) error {
	data := database.InsertCurrencyDataInDatabase(
		ctx,
		currencyData.USDBRL.Code,
		currencyData.USDBRL.Codein,
		currencyData.USDBRL.Name,
		currencyData.USDBRL.High,
		currencyData.USDBRL.Low,
		currencyData.USDBRL.VarBid,
		currencyData.USDBRL.PctChange,
		currencyData.USDBRL.Bid,
		currencyData.USDBRL.Ask,
		currencyData.USDBRL.Timestamp,
		currencyData.USDBRL.CreateDate)
	if data != nil {
		return data
	}
	return nil
}
