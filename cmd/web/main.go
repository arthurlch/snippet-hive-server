package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "4000", "HTTP Network Address")  
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Println("Starting server on port",*addr )
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
