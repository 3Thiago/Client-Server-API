package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() error {
	db, err := sql.Open("sqlite3", "currencyData.db")
	if err != nil {
		return err
	}
	defer db.Close()
	sts := `CREATE TABLE IF NOT EXISTS currencyData(id INTEGER PRIMARY KEY,
		Code        TEXT,
		Codein      TEXT,
		Name        TEXT,
		High        TEXT,
		Low         TEXT,
		VarBid      TEXT,
		PctChange   TEXT,
		Bid         TEXT,
		Ask         TEXT,
		Timestamp   TEXT,
		CreateDate  TEXT);`
	_, err = db.Exec(sts)
	if err != nil {
		return err
	}
	return nil
}

func InsertCurrencyDataInDatabase(
	ctx context.Context,
	Code string,
	Codein string,
	Name string,
	High string,
	Low string,
	VarBid string,
	PctChange string,
	Bid string,
	Ask string,
	Timestamp string,
	CreateDate string) error {

	db, err := sql.Open("sqlite3", "currencyData.db")
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}
	insertCurrencyDataSQL := `INSERT INTO currencyData(
		Code	   ,   
		Codein     ,
		Name       ,
		High       ,
		Low        ,
		VarBid     ,
		PctChange  ,
		Bid        ,
		Ask        ,
		Timestamp  ,
		CreateDate ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.PrepareContext(ctx, insertCurrencyDataSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.ExecContext(ctx, Code, Codein, Name, High, Low, VarBid, PctChange, Bid, Ask, Timestamp, CreateDate)
	if err != nil {
		return err
	}
	return nil
}
