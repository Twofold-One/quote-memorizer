package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/quote", app.showQuote)
	mux.HandleFunc("/quote/create", app.createQuote)
	
	fileServer := http.FileServer(restrictedFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/static", http.NotFoundHandler())

	return mux
}
