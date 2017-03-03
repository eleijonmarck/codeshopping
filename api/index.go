package api

import (
	"net/http"
)

// IndexHandler handles the first root handle and says PONG
func IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong"))
	})
}
