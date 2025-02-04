package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/", helloHandler)
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server starting on port 8080")
	// http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ciao Mario!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Io sono Mario!")
}
