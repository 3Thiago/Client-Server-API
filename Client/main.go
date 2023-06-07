package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	CreateFile(string(body))
}

func CreateFile(cotacao string) {
	filename := "cotacao.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	err := ioutil.WriteFile(filename, []byte("Dolar:"+cotacao), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
