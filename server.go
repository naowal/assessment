package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Expend struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
	Note   string `json:"note"`
	Tag    []string `json:"tag"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS expends ( id SERIAL PRIMARY KEY, title TEXT, amount INT, note TEXT, tag TEXT [] );	
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table ", err)
	}

	fmt.Println("create table success")
}
