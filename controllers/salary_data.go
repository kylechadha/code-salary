package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/kylechadha/code-salary/app"
	"github.com/kylechadha/code-salary/models"
)

// SalaryData Controller type.
type salaryDataController struct {
	databaseService app.IDatabaseService
}

// SalaryData Controller constructor function.
func NewSalaryDataController(app *app.Ioc) *salaryDataController {
	return &salaryDataController{app.DatabaseService}
}

// curl -XPOST -H 'Content-Type: application/json' -d '{"company": "Google", "city": "New York", "state": "New York", "Country": "USA", "Base": 136000, "Bonus": "20%", "Perks": 20000, "Stack": ["linux", "mysql", "apache", "php"]}' http://localhost:3000/api/salaryData
// SalaryDataCreate handler.
func (c *salaryDataController) SalaryDataCreate(w http.ResponseWriter, r *http.Request) (error, int) {

	// Create a new SalaryData struct and set the DateAdded.
	s := models.SalaryData{}
	s.DateAdded = time.Now()

	// Decode the JSON onto the struct.
	json.NewDecoder(r.Body).Decode(&s)

	// Create the item via the DB Service.
	err := c.databaseService.Create(s)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusCreated
}

// curl -H "Content-Type: application/json" http://localhost:3000/api/salaryData/3
func (c *salaryDataController) SalaryDataFind(w http.ResponseWriter, r *http.Request) (error, int) {

	// Get the salaryData ID from the route.
	vars := mux.Vars(r)
	sId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Retrieve the item from the DB Service.
	s, err := c.databaseService.Find(sId)
	if err != nil {
		// ** Need to check that error is what we see when it can't find the record
		return err, http.StatusNotFound
	}

	// Marshal the document as JSON.
	json, err := json.Marshal(s)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Write the JSON to the response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

	return nil, http.StatusOK
}

func (c *salaryDataController) SalaryDataFindN(w http.ResponseWriter, r *http.Request) (error, int) {

	// Get n, an optional sort, and an optional sort direction from the
	// query params.
	qp := r.URL.Query()

	qpN, ok := qp["n"]
	var n int
	var err error
	if ok {
		n, err = strconv.Atoi(qpN[0])
		if err != nil {
			return fmt.Errorf("Error parsing 'n' query param: %s", err), http.StatusInternalServerError
		}
	}

	qpSort, ok := qp["sort"]
	var sort string
	if ok {
		sort = strings.ToLower(qpSort[0])
	}

	qpAsc, ok := qp["asc"]
	var asc bool
	if ok {
		asc, err = strconv.ParseBool(qpAsc[0])
		if err != nil {
			return fmt.Errorf("Error parsing 'asc' query param: %s", err), http.StatusInternalServerError
		}
	}

	// Retrieve n results in the salaryData collection.
	ss, err := c.databaseService.FindN(n, sort, asc)
	if err != nil {
		return err, http.StatusNotFound
	}

	// Marshal the documents as JSON.
	json, err := json.Marshal(ss)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Write the JSON to the response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

	return nil, http.StatusOK
}

// func (c *salaryDataController) SalaryDataFindAll(w http.ResponseWriter, r *http.Request) (error, int) {

// 	// // Retrieve all documents in the salaryData collection.
// 	// salaryData, err := c.databaseService.FindAll("salaryData")
// 	// if err != nil {
// 	// 	return err, http.StatusNotFound
// 	// }

// 	// // Marshal the documents as JSON.
// 	// json, err := json.Marshal(salaryData)
// 	// if err != nil {
// 	// 	return err, http.StatusInternalServerError
// 	// }

// 	// // Write the JSON to the response.
// 	// w.Header().Set("Content-Type", "application/json")
// 	// w.Write(json)

// 	return nil, http.StatusOK
// }
