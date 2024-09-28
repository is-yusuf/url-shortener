package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/is-yusuf/url-shortener/pkg/urlshortener"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", urlshortener.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", urlshortener.RedirectURL).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
