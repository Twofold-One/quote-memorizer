package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Twofold-One/quote-memorizer/pkg/models/postgresql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	quotes *postgresql.QuoteModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":4000", "Network HTTP address")
	dsn := flag.String("dsn", "postgresql://user:pass@localhost:5432/quote-memorizer", "DB Source path")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := sql.Open("pgx", *dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to DB!")

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		quotes: &postgresql.QuoteModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Server is running on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

type restrictedFileSystem struct {
	fs http.FileSystem
}

func (rfs restrictedFileSystem) Open(path string) (http.File, error) {
	f, err := rfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := rfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}