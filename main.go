package main

import (
	"log"
	"net/http"
)

// home is handler function.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from quote-memorizer"))
}

func main() {
	// mux uses http.NewServeMux() method to initialize new router, 
	//  and registers home function as handler to URL template "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Server startup
	log.Println("Server is running on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}