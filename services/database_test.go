package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/kylechadha/code-salary/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	// Create a mock DB.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unable to create a mock DB, %s", err)
	}
	defer db.Close()

	// Define the mock expectations.
	mock.ExpectExec(`^INSERT INTO code_salary`).WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectExec(`^INSERT IGNORE INTO stack`).WithArgs("golang").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`^INSERT INTO salary_stack`).WithArgs(2, "golang").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`^INSERT IGNORE INTO stack`).WithArgs("docker").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`^INSERT INTO salary_stack`).WithArgs(2, "docker").WillReturnResult(sqlmock.NewResult(1, 1))

	// Instantiate the DB service and attempt to Create new salaryData.
	d := databaseService{db}
	s := models.SalaryData{
		Company:   "Bonobos",
		City:      "Los Angeles",
		State:     "California",
		Country:   "USA",
		Base:      146000,
		Bonus:     "20%",
		Perks:     15000,
		Stack:     []string{"golang", "docker"},
		DateAdded: time.Now(),
	}
	err = d.Create(s)

	// Test the results.
	if err != nil {
		t.Errorf("An error was returned from databaseService.Create: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

	// Create a mock DB.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unable to create a mock DB, %s", err)
	}
	defer db.Close()

	// Define the mock expectations.
	want := models.SalaryData{
		Id:        1,
		Company:   "Bonobos",
		City:      "Los Angeles",
		State:     "California",
		Country:   "USA",
		Base:      146000,
		Bonus:     "20%",
		Perks:     15000,
		Stack:     []string{"golang", "docker"},
		DateAdded: time.Now(),
	}
	rows := sqlmock.NewRows([]string{"id", "company", "city", "state", "country", "base", "bonus", "perks", "date_added", "stack"}).
		AddRow(want.Id, want.Company, want.City, want.State, want.Country, want.Base, want.Bonus, want.Perks, want.DateAdded, "golang docker")
	mock.ExpectQuery(`^SELECT (.+) FROM code_salary`).WithArgs(1).WillReturnRows(rows)

	// Instantiate the DB service and attempt to Find ID 1.
	d := databaseService{db}
	got, err := d.Find(1)

	// Test the results.
	if err != nil {
		t.Errorf("An error was returned from databaseService.Find: %s", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %+v to equal %+v", got, want)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}
}

func TestFindN(t *testing.T) {
	t.Parallel()

	// Create a mock DB.
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Unable to create a mock DB, %s", err)
	}
	defer db.Close()

	// Define the mock expectations.
	s1 := models.SalaryData{
		Id:        1,
		Company:   "Bonobos",
		City:      "Los Angeles",
		State:     "California",
		Country:   "USA",
		Base:      146000,
		Bonus:     "20%",
		Perks:     15000,
		Stack:     []string{"golang", "docker"},
		DateAdded: time.Now(),
	}
	s2 := models.SalaryData{
		Id:        2,
		Company:   "Google",
		City:      "New York",
		State:     "New York",
		Country:   "USA",
		Base:      175000,
		Bonus:     "10%",
		Perks:     25000,
		Stack:     []string{"c", "c++", "python"},
		DateAdded: time.Now(),
	}
	s3 := models.SalaryData{
		Id:        3,
		Company:   "Qualcomm",
		City:      "San Diego",
		State:     "California",
		Country:   "USA",
		Base:      116000,
		Bonus:     "35%",
		Perks:     10000,
		Stack:     []string{"perl", "mysql"},
		DateAdded: time.Now(),
	}

	want1 := []models.SalaryData{s1, s2, s3}
	rows1 := sqlmock.NewRows([]string{"id", "company", "city", "state", "country", "base", "bonus", "perks", "date_added", "stack"}).
		AddRow(s1.Id, s1.Company, s1.City, s1.State, s1.Country, s1.Base, s1.Bonus, s1.Perks, s1.DateAdded, "golang docker").
		AddRow(s2.Id, s2.Company, s2.City, s2.State, s1.Country, s2.Base, s2.Bonus, s2.Perks, s2.DateAdded, "c c++ python").
		AddRow(s3.Id, s3.Company, s3.City, s3.State, s3.Country, s3.Base, s3.Bonus, s3.Perks, s3.DateAdded, "perl mysql")
	mock.ExpectQuery(`^SELECT (.+) FROM code_salary`).WithArgs(1000).WillReturnRows(rows1)

	want2 := []models.SalaryData{s1, s2}
	rows2 := sqlmock.NewRows([]string{"id", "company", "city", "state", "country", "base", "bonus", "perks", "date_added", "stack"}).
		AddRow(s1.Id, s1.Company, s1.City, s1.State, s1.Country, s1.Base, s1.Bonus, s1.Perks, s1.DateAdded, "golang docker").
		AddRow(s2.Id, s2.Company, s2.City, s2.State, s2.Country, s2.Base, s2.Bonus, s2.Perks, s2.DateAdded, "c c++ python")
	mock.ExpectQuery(`^SELECT (.+) FROM code_salary`).WithArgs(2).WillReturnRows(rows2)

	// Instantiate the DB service and attempt to FindN, n=0 and n=2.
	d := databaseService{db}
	got1, err1 := d.FindN(0, "company", true)
	got2, err2 := d.FindN(2, "", false)

	// Test the results.
	if err1 != nil {
		t.Errorf("An error was returned from databaseService.FindN: %s", err1)
	}
	if err2 != nil {
		t.Errorf("An error was returned from databaseService.FindN: %s", err2)
	}
	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("Expected %+v to equal %+v", got1, want1)
	}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("Expected %+v to equal %+v", got1, want1)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}
}
