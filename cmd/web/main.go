package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"html/template"
	"log"
	"net/http"
	"os"
	"snippetbox.dany.net/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

var ctx = context.Background()

func main() {
	var err error

	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	connString := "postgres://postgres:Danalex4610@localhost:5432/snippetbox"
	conn, err := openDB(connString)

	if err != nil {
		errorLog.Fatalf("\nUnable to connection to database: %v\n", err)
	}
	defer conn.Close()

	infoLog.Println("Connected!")

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{
			DB: conn,
		},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(connString string) (*pgxpool.Pool, error) {

	db, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return db, err
}
