package main

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}


	w.Write([]byte("Hello from SnippetBox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet %d",id)
}

// func specificSpinnet(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Showing a specific snippet"))
// }

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return	
	}

	w.Write([]byte("Create a new snippet"))
}


func main(){
	//HTTP Request Multiplier(Router).Dispatches the incoming requests to the matching URL path
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view",snippetView)
	mux.HandleFunc("/snippet/create",createSnippet)


	log.Println("Server starting on 7001")
	err := http.ListenAndServe(":7001", mux)
	log.Fatal(err)
}