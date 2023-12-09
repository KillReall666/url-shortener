package main

import (
	"log"
	"net/http"

	"github.com/KillReall666/url-shortener/internal/handlers/addurl"
	"github.com/KillReall666/url-shortener/internal/handlers/geturl"
	"github.com/KillReall666/url-shortener/internal/storage"

	"github.com/gorilla/mux"
)

func main() {

	Store := storage.New()

	addURL := addurl.NewAddUrlHandler(Store)
	getURL := geturl.NewGetUrlHandler(Store)

	r := mux.NewRouter()
	r.HandleFunc("/", addURL.AddURL).Methods("POST")
	r.HandleFunc("/{id}", getURL.GetURL).Methods("GET")

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
	//СЮДА НЕ ПРОХОДИТ
	log.Println("server started on:", srv.Addr)
}
