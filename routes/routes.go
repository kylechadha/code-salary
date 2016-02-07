package routes

import (
	"net/http"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/utils"

	"github.com/gorilla/mux"
)

func NewRouter(app *app.Ioc) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// API routes.
	d := app.SalaryDataController
	salaryData := router.PathPrefix("/api").Subrouter()
	salaryData.Handle("/day", utils.AppHandler(d.SalaryDataCreate))
	salaryData.Handle("/day/{id}/", utils.AppHandler(d.SalaryDataFind))
	salaryData.Handle("/salaryData/", utils.AppHandler(d.SalaryDataFindAll))

	// Static files.
	router.PathPrefix("/libs").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/scripts").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/styles").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/views")))

	// Angular routes...

	return router
}
