package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("started-service")

	r := mux.NewRouter()
	r.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST", "OPTIONS")
	r.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8080", r))
	//
	//fmt.Println("started-service")
	//http.HandleFunc("/upload", uploadHandler)
	//log.Fatal(http.ListenAndServe(":8080", nil))

}
