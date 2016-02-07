package utils

import (
	"net/http"
	"strings"
)

// RestrictDir Helper.
// Helper to restrict access to the directory tree listing.
func RestrictDir(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
