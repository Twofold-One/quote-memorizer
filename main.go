package main

import (
	"fmt"
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
	w.Write([]byte("Hello from quote-memorizer"))
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

func main() {
	// mux uses http.NewServeMux() method to initialize new router, 
	//  and registers home function as handler to URL template "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/quote", showQuote)
	mux.HandleFunc("/quote/create", createQuote)

	// Server startup
	log.Println("Server is running on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}