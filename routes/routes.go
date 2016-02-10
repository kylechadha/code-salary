package routes

import (
	"net/http"

	"github.com/kylechadha/code-salary/app"
	"github.com/kylechadha/code-salary/utils"

	"github.com/gorilla/mux"
)

// ToDos...
// - Move AppHandler and RestrictDir here.

func NewRouter(app *app.Ioc) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// API routes.
	d := app.SalaryDataController
	salaryData := router.PathPrefix("/api").Subrouter()
	salaryData.Handle("/salaryData/{id}/", utils.AppHandler(d.SalaryDataFind))
	salaryData.Handle("/salaryData", utils.AppHandler(d.SalaryDataCreate))   // post
	salaryData.Handle("/salaryData/", utils.AppHandler(d.SalaryDataFindAll)) // get

	// Static files.
	router.PathPrefix("/libs").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/scripts").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/styles").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/views")))

	// Angular routes...

	return router
}
