package main

import (
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from SnippetBox"))
}


func main(){
	//HTTP Request Multiplier(Router).Dispatches the incoming requests to the matching URL path
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)

	log.Println("Server starting on 7001")
	err := http.ListenAndServe(":7001", mux)
	log.Fatal(err)
}