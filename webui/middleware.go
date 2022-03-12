package webui

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Interface("Header", r.Header).Msg("")
		log.Info().Str("uri", r.RequestURI).Str("method", r.Method).Msg("")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
