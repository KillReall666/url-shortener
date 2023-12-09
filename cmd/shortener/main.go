package main

import (
	"github.com/KillReall666/url-shortener/url-shortener/internal/handlers/addurl"
	"github.com/KillReall666/url-shortener/url-shortener/internal/handlers/geturl"
	"github.com/KillReall666/url-shortener/url-shortener/internal/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	Store := storage.New()

	addUrl := addurl.NewAddUrlHandler(Store)
	getUrl := geturl.NewGetUrlHandler(Store)

	r := mux.NewRouter()
	r.HandleFunc("/", addUrl.AddUrl).Methods("POST")
	r.HandleFunc("/{id}", getUrl.GetUrl).Methods("GET")

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
