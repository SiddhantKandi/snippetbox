package main

import (
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}


	w.Write([]byte("Hello from SnippetBox"))
}

func specificSpinnet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Showing a specific snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		w.Header().Set("ALLOW", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return	
	}

	w.Write([]byte("Create a new snippet"))
}


func main(){
	//HTTP Request Multiplier(Router).Dispatches the incoming requests to the matching URL path
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view",specificSpinnet)
	mux.HandleFunc("/snippet/create",createSnippet)


	log.Println("Server starting on 7001")
	err := http.ListenAndServe(":7001", mux)
	log.Fatal(err)
}