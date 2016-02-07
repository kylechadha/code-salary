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
	d := app.DaysController
	days := router.PathPrefix("/api").Subrouter()
	days.Handle("/day", utils.AppHandler(d.DaysCreate))
	days.Handle("/day/{id}/", utils.AppHandler(d.DaysFind))
	days.Handle("/days/", utils.AppHandler(d.DaysFindAll))

	// Static files.
	router.PathPrefix("/libs").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/scripts").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/styles").Handler(utils.RestrictDir(http.FileServer(http.Dir("./public/"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/views")))

	// Angular routes...

	return router
}
