package services

import (
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

	// Define the mock expectations.
	mock.ExpectExec(`^INSERT INTO code_salary`).WillReturnResult(sqlmock.NewResult(2, 1)) // change this to two and should fail
	mock.ExpectExec(`^INSERT IGNORE INTO stack`).WithArgs("golang").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`^INSERT INTO salary_stack`).WithArgs(2, "golang").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`^INSERT IGNORE INTO stack`).WithArgs("docker").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`^INSERT INTO salary_stack`).WithArgs(2, "docker").WillReturnResult(sqlmock.NewResult(1, 1))

	// Instantiate the DB service with the mock.
	d := databaseService{db}

	// Create new salaryData and hit Create.
	s := models.SalaryData{
		Company: "Bonobos",
		City:    "Los Angeles",
		State:   "California",
		Country: "USA",
		Base:    146000,
		Bonus:   "20%",
		Perks:   25000,
		Stack:   []string{"golang", "docker"},
	}
	s.DateAdded = time.Now()
	err = d.Create(s)

	// Check the results.
	if err != nil {
		t.Errorf("An error was returned from databaseService.Create: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}

}
