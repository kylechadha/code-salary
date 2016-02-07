package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// curl -XPOST -H 'Content-Type: application/json' -d '{"salaryDataOfTheWeek": "SatursalaryData", "salaryDataOfTheWeekType": "Weekend", "successfulWakeUp": true, "morningWork": true, "morningWorkType": "omnia app", "workedOut": true, "workedOutType": "swam", "plannedNextSalaryData": true}' http://localhost:3000/api/salaryData
// SalaryDataCreate handler.
func (c *salaryDataController) SalaryDataCreate(w http.ResponseWriter, r *http.Request) (error, int) {

	// Create a new SalaryData struct and set the ObjectId.
	salaryData := models.SalaryData{}

	// Decode the JSON onto the struct.
	json.NewDecoder(r.Body).Decode(&salaryData)

	// Create the SalaryData via the Database Service.
	err := c.databaseService.Create(salaryData)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusCreated
}

func (c *salaryDataController) SalaryDataFind(w http.ResponseWriter, r *http.Request) (error, int) {

	// Get the salaryData ID from the params.
	vars := mux.Vars(r)
	salaryDataId, err := strconv.Atoi(vars["id"])
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Retrieve the document.
	salaryData, err := c.databaseService.Find(salaryDataId)
	if err != nil {
		// ** Need to check that error is what we see when it can't find the record
		return err, http.StatusNotFound
	}

	// Marshal the document as JSON.
	json, err := json.Marshal(salaryData)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	// Write the JSON to the response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

	return nil, http.StatusOK
}

func (c *salaryDataController) SalaryDataFindAll(w http.ResponseWriter, r *http.Request) (error, int) {

	// // Retrieve all documents in the salaryData collection.
	// salaryData, err := c.databaseService.FindAll("salaryData")
	// if err != nil {
	// 	return err, http.StatusNotFound
	// }

	// // Marshal the documents as JSON.
	// json, err := json.Marshal(salaryData)
	// if err != nil {
	// 	return err, http.StatusInternalServerError
	// }

	// // Write the JSON to the response.
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(json)

	return nil, http.StatusOK
}
