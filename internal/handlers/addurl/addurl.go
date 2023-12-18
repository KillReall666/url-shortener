package addurl

import (
	"github.com/KillReall666/url-shortener/internal/config"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/KillReall666/url-shortener/internal/storage"

	"github.com/teris-io/shortid"
)

type AddURLHandler struct {
	Storage *storage.Storage
	cfg     config.RunConfig
}

func NewAddURLHandler(s *storage.Storage, conf config.RunConfig) *AddURLHandler {
	return &AddURLHandler{
		Storage: s,
		cfg:     conf,
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
	w.Write([]byte(a.cfg.ShortURLAddress + id))

}
