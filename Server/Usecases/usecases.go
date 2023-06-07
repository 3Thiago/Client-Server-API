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
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var currencyData model.CurrencyData
	err = json.Unmarshal(body, &currencyData)
	if err != nil {
		return nil, err
	}
	saveDataCtx(ctx, &currencyData)
	return &currencyData, nil
}

func saveDataCtx(ctx context.Context, currencyData *model.CurrencyData) {
	select {
	case <-ctx.Done():
		//fmt.Println("API call is canceled...")
		return
	case <-time.After(10 * time.Millisecond):
		//fmt.Println("Successfully saved data...")
		InsertCurrencyData(ctx, currencyData)
	}
}

func InsertCurrencyData(ctx context.Context, currencyData *model.CurrencyData) error {
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
