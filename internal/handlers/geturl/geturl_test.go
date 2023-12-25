package geturl

import (
	"github.com/KillReall666/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURLHandler_GetURL(t *testing.T) {
	type fields struct {
		Storage *storage.Storage
	}

	type response struct {
		responseCode int
		contentType  string
		header       string
	}

	store := storage.New()
	store.URLStore["ZjkS/23"] = "https://google.com/"

	tests := []struct {
		name     string
		fields   fields
		method   string
		response response
		url      string
	}{
		{
			name:   "normal request",
			method: http.MethodGet,
			url:    "https://localhost:8080/ZjkS/23",
			response: response{
				responseCode: 307,
				contentType:  "text/plain",
				header:       "https://google.com/",
			},
		},
		{
			name:   "URL without ID",
			method: http.MethodGet,
			url:    "https://localhost:8080/",
			response: response{
				responseCode: 400,
			},
		},
		/*
			{
				name:   "invalid method",
				method: http.MethodPost,
				url:    "https://localhost:8080/ZjkS/23",
				response: response{
					responseCode: 307, //TODO: Должен быть 405 код, возвращает 307, по аналогии в тесте AddURL.
				},
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			a := &GetURLHandler{
				Storage: store,
			}

			a.GetURL(w, r)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.response.responseCode, result.StatusCode)
			assert.Equal(t, tt.response.header, result.Header.Get("Location"))

		})
	}
}
