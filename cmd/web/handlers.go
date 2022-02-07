package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// home is main page handler function.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// showQuote is quote show handler function.
func showQuote(w http.ResponseWriter, r *http.Request) {
	// get id from request, convert it to int and check if
	// value less than 1 response 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Show chosen quote with ID %d...\n", id)
}

// createQuote is quote creation handler function.
func createQuote(w http.ResponseWriter, r *http.Request) {


	// r.Method checks if request is POST.
	if r.Method != http.MethodPost {

		// Header().Set() adds "Allow: POST" to the header
		w.Header().Set("Allow", http.MethodPost)
		// http.Error sends status code and a message
		http.Error(w, "Method is restricted", http.StatusMethodNotAllowed)
		return 
	}
	w.Write([]byte("Form to create new quote"))
}