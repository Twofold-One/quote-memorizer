package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/quote", showQuote)
	mux.HandleFunc("/quote/create", createQuote)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/static", http.NotFoundHandler())
	

	log.Println("Server is running on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}


// TODO Fix "type restrictedFileSystem is unused (U1000)"
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