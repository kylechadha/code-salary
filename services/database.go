package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kylechadha/code-salary/app"
	"github.com/kylechadha/code-salary/models"

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
		log.Fatal(err)
	}

	dbPassword, err := app.ConfigService.GetConfig("db_password")
	if err != nil {
		fmt.Println("Config file does not include a 'db_password'.")
		log.Fatal(err)
	}

	dbHost, err := app.ConfigService.GetConfig("db_host")
	if err != nil {
		fmt.Println("Config file does not include a 'db_host'.")
		log.Fatal(err)
	}

	dbName, err := app.ConfigService.GetConfig("db_name")
	if err != nil {
		fmt.Println("Config file does not include a 'db_name'.")
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@"+dbHost+"/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Temporary ... rather, we'll want to close on shutdown.

	// Make sure the database credentials are valid and a connection is possible.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &databaseService{db: db}
}

func (d *databaseService) Create(data models.SalaryData) error {

	result, err := d.db.Exec("INSERT INTO code_salary (company, city, state, country, base, bonus, perks, date_added) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", data)
	if err != nil {
		return err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	rowCnt, err := result.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("TESTING!!! ID = %d, affected = %d\n", lastId, rowCnt)

	return nil
}

func (d *databaseService) Find(id int) (models.SalaryData, error) {

	var salaryData models.SalaryData

	// ** What if instead of doing this off hand here, we prepare a statement in the constructor and this just calls that? Any advantage?
	err := d.db.QueryRow("SELECT id, company FROM code_salary WHERE id = ?", id).Scan(&salaryData)
	if err != nil {
		return salaryData, err
	}

	// rows, err := d.db.Query("SELECT id, company FROM code_salary WHERE id = ?", id)
	// if err != nil {
	// 	return salaryData, err
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	err := rows.Scan(&salaryData)
	// 	if err != nil {
	// 		return salaryData, err
	// 	}
	// }
	// err = rows.Err()
	// if err != nil {
	// 	return salaryData, err
	// }

	return salaryData, err
}

func (d *databaseService) FindAll(collection string) ([]interface{}, error) {

	// Find all documents in the collection.

	return nil, nil
}
