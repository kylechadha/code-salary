package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
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

	var s models.SalaryData
	var stack string

	// ** What if instead of doing this off hand here, we prepare a statement in the constructor and this just calls that? Any advantage?
	err := d.db.QueryRow(`
		SELECT id, company, city, state, country, base, bonus, perks, date_added, IFNULL(group_concat(DISTINCT s.stack_name SEPARATOR ' '), '')
		FROM code_salary AS c
		INNER JOIN salary_stack AS s
		WHERE c.id = s.salarydata_id
		AND c.id = ?
		GROUP BY c.id
		`, id).Scan(&s.Id, &s.Company, &s.City, &s.State, &s.Country, &s.Base, &s.Bonus, &s.Perks, &s.DateAdded, &stack)
	if err != nil {
		return s, err
	}

	// Converting stack (string returned from the db join) to a proper
	// slice to be stored in the models.salaryData type. If we weren't attached
	// to mysql, this wouldn't have to be so clumsy.
	if stack != "" {
		s.Stack = strings.Split(stack, " ")
	}

	return s, err
}

func (d *databaseService) FindN(n int, sort string, asc bool) ([]models.SalaryData, error) {

	// Set n to be 1000 when no limit is specified.
	if n == 0 {
		n = 1000
	}

	// Validate the sort field.
	qFields := map[string]bool{
		"company":   true,
		"city":      true,
		"state":     true,
		"country":   true,
		"base":      true,
		"bonus":     true,
		"perks":     true,
		"dateAdded": true,
	}
	if ok := qFields[sort]; !ok {
		sort = "null"
	}

	// Just in case there's some sql injection voodoo, let's double check 'sort'.
	valid, err := regexp.Compile("^[A-Za-z0-9_]+$")
	if err != nil {
		return nil, err
	}
	if !valid.MatchString(sort) {
		return nil, errors.New("invalid sort query parameter")
	}

	// Convert the asc bool to the corresponding string.
	var ascStr string
	if asc {
		ascStr = "ASC"
	} else {
		ascStr = "DESC"
	}

	// Construct the SQL query.
	query := fmt.Sprintf(`
		SELECT id, company, city, state, country, base, bonus, perks, date_added, IFNULL(group_concat(DISTINCT s.stack_name SEPARATOR ' '), '')
		FROM code_salary AS c
		INNER JOIN salary_stack AS s
		WHERE c.id = s.salarydata_id
		GROUP BY c.id
		ORDER BY %s %s
		LIMIT ?
		`, sort, ascStr)

	rows, err := d.db.Query(query, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ss []models.SalaryData
	for rows.Next() {
		var s models.SalaryData
		var stack string
		err := rows.Scan(&s.Id, &s.Company, &s.City, &s.State, &s.Country, &s.Base, &s.Bonus, &s.Perks, &s.DateAdded, &stack)
		if err != nil {
			return nil, err
		}

		if stack != "" {
			s.Stack = strings.Split(stack, " ")
		}

		ss = append(ss, s)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return ss, nil
}

// * How / when to add indexes?

// For Tom:
// - write these methods
// - write the tests
