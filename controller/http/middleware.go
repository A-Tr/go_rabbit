package http

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func RequestIdMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		requestId := uuid.NewV4()
		r.Header.Set("request-id", requestId.String())
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
