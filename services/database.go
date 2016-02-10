package services

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/kylechadha/code-salary/app"
	"github.com/kylechadha/code-salary/models"

	_ "github.com/go-sql-driver/mysql"
)

// Database Service type.
// MySQL probably isn't the best DB for our purposes,
// but we're going to use it any way as a learning
// exercise :)
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

	// 'parseTime=true' is an optional flag for the go-sql-driver,
	// to add Scan() support for parsing into time.Time.
	// https://github.com/go-sql-driver/mysql/issues/9
	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@"+dbHost+"/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Make sure the database credentials are valid and a connection is possible.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &databaseService{db: db}
}

func (d *databaseService) Create(data models.SalaryData) error {

	result, err := d.db.Exec(`INSERT INTO code_salary (company, city, state, country, base, bonus, perks, date_added) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, data.Company, data.City, data.State, data.Country, data.Base, data.Bonus, data.Perks, data.DateAdded)
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
	log.Printf("TESTING! ID = %d, affected = %d\n", lastId, rowCnt)

	// Note: You could also read the `stack` table in the constructor and create a cache to avoid hitting the DB.
	for _, elem := range data.Stack {
		stack := strings.ToLower(elem)

		// I wonder if you can insert multiple values in one query...
		_, err := d.db.Exec("INSERT IGNORE INTO stack VALUES (?)", stack)
		if err != nil {
			return err
		}

		_, err = d.db.Exec("INSERT INTO salary_stack (salarydata_id, stack_name) VALUES (?, ?)", lastId, stack)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *databaseService) Find(id int) (models.SalaryData, error) {

	var salaryData models.SalaryData

	// ** What if instead of doing this off hand here, we prepare a statement in the constructor and this just calls that? Any advantage?
	// ** Could simplify this with GORP or sqlstruct pkgs... might not be worth the benefit though
	// `SELECT p.*, f.*
	// FROM person p
	// INNER JOIN person_fruit pf
	// ON pf.person_id = p.id
	// INNER JOIN fruits f
	// ON f.fruit_name = pf.fruit_name`
	// err := d.db.QueryRow("SELECT ... FROM ", id).Scan(&salaryData.Id, &salaryData.Company, &salaryData.City, &salaryData.State, &salaryData.Country, &salaryData.Base, &salaryData.Bonus, &salaryData.Perks, &salaryData.DateAdded)
	err := d.db.QueryRow("SELECT id, company, city, state, country, base, bonus, perks, date_added FROM code_salary WHERE id = ?", id).Scan(&salaryData.Id, &salaryData.Company, &salaryData.City, &salaryData.State, &salaryData.Country, &salaryData.Base, &salaryData.Bonus, &salaryData.Perks, &salaryData.DateAdded)
	if err != nil {
		return salaryData, err
	}

	return salaryData, err
}

func (d *databaseService) FindAll(collection string) ([]interface{}, error) {

	// Find all documents in the collection.

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

	return nil, nil
}
