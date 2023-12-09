package addurl

import (
	"io"
	"log"
	"net/http"

	"github.com/KillReall666/url-shortener/internal/storage"

	"github.com/teris-io/shortid"
)

type AddURLHandler struct {
	Storage *storage.Storage
}

func NewAddUrlHandler(s *storage.Storage) *AddURLHandler {
	return &AddURLHandler{
		Storage: s,
	}
}

func (a *AddURLHandler) AddURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	url := string(body)
	id, err := shortid.Generate()
	if err != nil {
		log.Println(err)
	}

	a.Storage.URLStore[id] = url

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + id))
}
