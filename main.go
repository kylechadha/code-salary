package main

import (
	"os"

	"github.com/kylechadha/omnia-app/app"
	"github.com/kylechadha/omnia-app/controllers"
	"github.com/kylechadha/omnia-app/routes"
	"github.com/kylechadha/omnia-app/services"

	"github.com/codegangsta/negroni"
)

func main() {

	// Config
	// ----------------------------

	// Set the port.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Object Graph
	// ----------------------------
	app := app.Ioc{}
	app.ConfigService = services.NewConfigService()
	app.DatabaseService = services.NewDatabaseService(&app)
	app.SalaryDataController = controllers.NewSalaryDataController(&app)

	// Router
	// ----------------------------
	router := routes.NewRouter(&app)

	// Do we want to use negroni again? Meh ... what about another way to do recovery
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(router)

	// Server
	// ----------------------------
	n.Run(":" + port)

}
