package main

import (
	"net/http"

	database "github.com/3Thiago/Client-Server-API/Server/Database"
	handler "github.com/3Thiago/Client-Server-API/Server/Handler"
)

func main() {
	database.CreateTable()
	http.HandleFunc("/cotacao", handler.GetPricepHandler)
	http.ListenAndServe(":8080", nil)
	

}
