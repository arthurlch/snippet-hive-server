package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet creation"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet creation"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
