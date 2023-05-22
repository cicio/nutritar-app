package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN string
	Domain string
	DB *sql.DB
}

func main() {
	// set application config
	var app application 


	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres sslmode=disable dbname=foods timezone=UTC connect_timeout=5", "Postgres connection string");
	flag.Parse();


	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn
	// close the Db connection as soon as the function finishes
	defer app.DB.Close()


	app.Domain = "nutrintel.com"

	log.Println("Starting Application on port", port)




	//start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}



}