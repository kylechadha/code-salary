package utils

import "net/http"

// AppHandler Helper.
// Helper used to wrap handlers and manage errors and status codes surfaced by the controller.
func AppHandler(a func(w http.ResponseWriter, r *http.Request) (error, int)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err, code := a(w, r); err != nil {
			http.Error(w, err.Error(), code)
		}
	})
}
