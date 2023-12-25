package geturl

import (
	"net/http"
	"strings"

	"github.com/KillReall666/url-shortener/internal/storage"
)

type GetURLHandler struct {
	Storage *storage.Storage
}

func NewGetURLHandler(s *storage.Storage) *GetURLHandler {
	return &GetURLHandler{
		Storage: s,
	}
}

func (g *GetURLHandler) GetURL(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	id := strings.TrimLeft(url, "/")

	respURL := g.Storage.URLStore[id]
	if respURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", respURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
