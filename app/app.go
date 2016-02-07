package app

import (
	"net/http"

	"github.com/kylechadha/code-salary/models"
)

type Ioc struct {
	ConfigService        IConfigService
	DatabaseService      IDatabaseService
	SalaryDataController ISalaryDataController
}

type IConfigService interface {
	GetConfig(key string) (string, error)
}

type IDatabaseService interface {
	Create(data models.SalaryData) error
	Find(id int) (models.SalaryData, error)
	FindAll(collection string) ([]interface{}, error)
}

type ISalaryDataController interface {
	SalaryDataCreate(w http.ResponseWriter, r *http.Request) (error, int)
	SalaryDataFind(w http.ResponseWriter, r *http.Request) (error, int)
	SalaryDataFindAll(w http.ResponseWriter, r *http.Request) (error, int)
}
