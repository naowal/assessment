package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("postgres", "irtiqnti:I1wJ5UtoD9us3QWT5kHic5DCvwdvSyIj@rosie.db.elephantsql.com/irtiqnt")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	log.Println("okay")
}
