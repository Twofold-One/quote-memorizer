package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home is main page handler function.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// showQuote is quote show handler function.
func (app *application) showQuote(w http.ResponseWriter, r *http.Request) {
	// get id from request, convert it to int and check if
	// value less than 1 response 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	fmt.Fprintf(w, "Show chosen quote with ID %d...\n", id)
}

// createQuote is quote creation handler function.
func (app *application) createQuote(w http.ResponseWriter, r *http.Request) {


	// r.Method checks if request is POST.
	if r.Method != http.MethodPost {

		// Header().Set() adds "Allow: POST" to the header
		w.Header().Set("Allow", http.MethodPost)
		// http.Error sends status code and a message
		app.clientError(w, http.StatusMethodNotAllowed)
		return 
	}
	w.Write([]byte("Form to create new quote"))
}