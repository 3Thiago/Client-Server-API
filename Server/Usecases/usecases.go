package usecases

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	database "github.com/3Thiago/Client-Server-API/Server/Database"
	model "github.com/3Thiago/Client-Server-API/Server/Model"
)

func GetCurrencyData() (*model.CurrencyData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Print("Error on create request in server: ", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error on get response in server: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var currencyData model.CurrencyData
	err = json.Unmarshal(body, &currencyData)
	if err != nil {
		log.Print("Error on unmarshal json data: ", err)
		return nil, err
	}
	insertCurrencyData(&currencyData)
	return &currencyData, nil
}

func insertCurrencyData(currencyData *model.CurrencyData) error {
	data := database.InsertCurrencyDataInDatabase(
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
