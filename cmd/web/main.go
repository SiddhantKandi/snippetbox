package main

import (
"log"
"net/http"
)


func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/view",snippetView)

	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :7001")

	err := http.ListenAndServe(":7001", mux)
	
	log.Fatal(err)
}