package main

import (
	"fmt"
	"log"
	"net/http"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	_, err := fmt.Fprintf(w, "You queried: %s", query)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/search", queryHandler)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
