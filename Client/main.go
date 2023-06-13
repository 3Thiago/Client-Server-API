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
	for {
		getPrice()
		time.Sleep(3 * time.Second)
	}

}
func getPrice() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Print("Error on create request: ", err)
	} else {
		log.Print("Request created with success!")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error on get response: ", err)
	} else {
		log.Print("Response received with success!")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	createFile(string(body))

}

func createFile(cotacao string) {
	filename := "cotacao.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Print("Error on create file: ", err)
		} else {
			log.Print("File created with success!")
		}
		defer file.Close()
	}
	err := ioutil.WriteFile(filename, []byte("Dolar:"+cotacao), 0644)
	if err != nil {
		log.Print("Error on write file: ", err)
	} else {
		log.Print("File writed with success!")
	}
}
