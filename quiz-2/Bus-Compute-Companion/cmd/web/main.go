// main file

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ACuellarbz/3162/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	user_info models.UserModel
}

func main() {
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("COMPUTE_DB_DSN"), "PostgreSQL DSN (Data Source Name)")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}
	app := &application{
		user_info: models.UserModel{DB: db},
	}

	defer db.Close()
	log.Println("Database connection pool established")

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	//use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //specify timeout
	defer cancel()                                                          //+Defer attached to a function and executes as the last thing
	//lets ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
