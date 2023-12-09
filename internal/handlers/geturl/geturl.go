package geturl

import (
	"net/http"
	"strings"

	"github.com/KillReall666/url-shortener/internal/storage"
)

type GetUrlHandler struct {
	Storage *storage.Storage
}

func NewGetUrlHandler(s *storage.Storage) *GetUrlHandler {
	return &GetUrlHandler{
		Storage: s,
	}
}

func (g *GetUrlHandler) GetUrl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	id := strings.TrimLeft(url, "/")

	respUrl := g.Storage.UrlStore[id]
	if respUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", respUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
