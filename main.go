package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/dtynn/interview/db"
	"github.com/dtynn/interview/handler"
)

func main() {
	fs := flag.NewFlagSet("app", flag.ExitOnError)

	var dsn string
	fs.StringVar(&dsn, "dsn", "", "database connection")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}

	if dsn == "" {
		log.Fatalln("dsn required")
	}

	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	db.Setup(conn)

	mux := http.NewServeMux()
	mux.HandleFunc("/teams/add", handler.AddTeam)
	mux.HandleFunc("/schedule", handler.Schedule)

	if err := http.ListenAndServe(":10080", mux); err != nil {
		log.Fatalln(err)
	}

	log.Println("shutdown")
}
