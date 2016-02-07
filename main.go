package main

import (
	"os"

	"github.com/kylechadha/code-salary/app"
	"github.com/kylechadha/code-salary/controllers"
	"github.com/kylechadha/code-salary/routes"
	"github.com/kylechadha/code-salary/services"

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
