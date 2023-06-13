package main

import (
	"log"
	"net/http"

	database "github.com/3Thiago/Client-Server-API/Server/Database"
	handler "github.com/3Thiago/Client-Server-API/Server/Handler"
)

func main() {
	database.CreateTable()
	log.Print("Server are running on port: 8080")
	http.HandleFunc("/cotacao", handler.GetPricepHandler)
	http.ListenAndServe(":8080", nil)
	

}
