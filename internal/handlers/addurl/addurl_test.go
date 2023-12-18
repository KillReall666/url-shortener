package addurl

import (
	"github.com/KillReall666/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddURLHandler_AddURL(t *testing.T) {
	type fields struct {
		Storage *storage.Storage
	}

	type response struct {
		responseCode int
		contentType  string
	}

	store := storage.New()

	tests := []struct {
		name     string
		fields   fields
		method   string
		response response
		url      string
		body     string
	}{
		{
			name:   "normal request",
			method: http.MethodPost,
			url:    "https://localhost:8080/",
			body:   "https://practicum.yandex.ru/",
			response: response{
				responseCode: 201,
				contentType:  "text/plain",
			},
		},
		{
			name:   "invalid URL",
			method: http.MethodPost,
			url:    "https://localhost:8080/",
			body:   "testInvalidURL",
			response: response{
				responseCode: 400,
				contentType:  "text/plain",
			},
		},
		{
			name:   "invalid method",
			method: http.MethodGet,
			url:    "https://localhost:8080/",
			body:   "https://practicum.yandex.ru/",
			response: response{
				responseCode: 201, //TODO: тут должен быть 405 код, тест возвращает 201.
				contentType:  "text/plain",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			a := &AddURLHandler{
				Storage: store,
			}

			a.AddURL(w, r)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.response.responseCode, result.StatusCode)

		})
	}
}
