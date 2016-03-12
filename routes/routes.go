package routes

import (
	"net/http"
	"strings"

	"github.com/kylechadha/code-salary/app"

	"github.com/gorilla/mux"
)

func NewRouter(app *app.Ioc) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// API routes.
	d := app.SalaryDataController
	salaryData := router.PathPrefix("/api").Subrouter()
	salaryData.Handle("/salaryData/{id}", appHandler(d.SalaryDataFind))
	salaryData.Handle("/salaryData", appHandler(d.SalaryDataFindN)).Methods("GET")
	salaryData.Handle("/salaryData", appHandler(d.SalaryDataCreate)).Methods("POST")

	// Static files.
	router.PathPrefix("/libs").Handler(restrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/scripts").Handler(restrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/styles").Handler(restrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/views").Handler(http.FileServer(http.Dir("./public")))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/views")))

	return router
}

// appHandler is used to wrap handlers and manage errors and status
// codes surfaced by the controller.
func appHandler(a func(w http.ResponseWriter, r *http.Request) (error, int)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err, code := a(w, r); err != nil {
			http.Error(w, err.Error(), code)
		}
	})
}

// restrictDir is used to restrict access to the directory tree listing.
func restrictDir(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
