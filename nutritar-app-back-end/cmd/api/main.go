package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = 8080

type application struct {
	DSN 			string
	Domain 			string
	DB 				repository.DatabaseRepo
	auth 			Auth
	JWTSecret 		string
	JWTIssuer 		string
	JWTAudience		string
	CookieDomain	string 
	
}

func main() {
	// set application config
	var app application 


	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres sslmode=disable dbname=foods timezone=UTC connect_timeout=5", "Postgres connection string");
	flag.StringVar(&app.JWTSecret, "jwt-secret","cicio123!","signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issues","nutrintel.com","signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience","nutrintel.com","signing audience")
	flag.StringVar(&app.CookieDomain, "cookie_domain","localhost","cookie_domain")
	flag.StringVar(&app.Domain, "domain","nutrintel.com","domain")
	flag.Parse();


	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	// close the DB connection as soon as the function finishes (defer keyword guarantees the connection only closes after the function refered by defer as finished procesing)
	// Recommended to avoid DB leaks
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer: app.JWTIssuer,
		Audience: app.JWTAudience,
		Secret: app.JWTSecret,
		TokenExpiry: time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath: "/",
		CookieName: "__Host-refresh_token",
		CookieDomain: app.CookieDomain,

	}


	//app.Domain = "nutrintel.com"

	log.Println("Starting Application on port", port)




	//start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}



}