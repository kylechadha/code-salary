package services

import (
	"database/sql"
	"fmt"

	"github.com/kylechadha/code-salary/app"

	_ "github.com/go-sql-driver/mysql"
)

// Database Service type.
type databaseService struct {
	db *sql.DB
}

// Database Service constructor function.
func NewDatabaseService(app *app.Ioc) *databaseService {
	dbUser, err := app.ConfigService.GetConfig("db_user")
	if err != nil {
		fmt.Println("Config file does not include a 'db_user'.")
		panic(err)
	}

	dbPassword, err := app.ConfigService.GetConfig("db_password")
	if err != nil {
		fmt.Println("Config file does not include a 'db_password'.")
		panic(err)
	}

	dbName, err := app.ConfigService.GetConfig("db_name")
	if err != nil {
		fmt.Println("Config file does not include a 'db_name'.")
		panic(err)
	}

	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return &databaseService{db: db}
}

// func (d *databaseService) Create(collection string, document interface{}) error {

// 	// Write the document to the database.

// 	return nil
// }

// func (d *databaseService) Find(collection string, oId bson.ObjectId, document interface{}) (interface{}, error) {

// 	// Find the document by ID.

// 	return nil, nil
// }

// func (d *databaseService) FindAll(collection string) ([]interface{}, error) {

// 	// Find all documents in the collection.

// 	return nil, nil
// }
