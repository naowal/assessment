package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Expense struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Amount int    `json:"amount"`
	Note   string `json:"note"`
	Tags    []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

func createExpandHandler(c echo.Context) error {
	e := Expense{}
	err := c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4::text[])  RETURNING id", e.Title, e.Amount, e.Note, pq.Array(e.Tags))
	err = row.Scan(&e.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, e)
}


var db *sql.DB

func main() {
	db, err := sql.Open("postgres", "postgres://irtiqnti:I1wJ5UtoD9us3QWT5kHic5DCvwdvSyIj@rosie.db.elephantsql.com/irtiqnti")
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS expenses ( id SERIAL PRIMARY KEY, title TEXT, amount INT, note TEXT, tags TEXT [] );	
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table ", err)
	}

	fmt.Println("create table success")

	e := echo.New()

	e.POST("/expenses", createExpandHandler)

	log.Fatal(e.Start(":"+ os.Getenv("PORT")))
}
