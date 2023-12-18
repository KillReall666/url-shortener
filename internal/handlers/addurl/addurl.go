package addurl

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/KillReall666/url-shortener/internal/storage"

	"github.com/teris-io/shortid"
)

type AddURLHandler struct {
	Storage *storage.Storage
}

func NewAddURLHandler(s *storage.Storage) *AddURLHandler {
	return &AddURLHandler{
		Storage: s,
	}
}

func (a *AddURLHandler) AddURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	URL := string(body)

	_, err = url.ParseRequestURI(URL)
	if err != nil {
		log.Println("invalid URL")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := shortid.Generate()
	if err != nil {
		log.Println(err)
	}

	a.Storage.URLStore[id] = URL

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + id))

}
