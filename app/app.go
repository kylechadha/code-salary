package app

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"
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
	Create(collection string, data interface{}) error
	Find(collection string, oId bson.ObjectId, document interface{}) (interface{}, error)
	FindAll(collection string) ([]interface{}, error)
}

type ISalaryDataController interface {
	SalaryDataCreate(w http.ResponseWriter, r *http.Request) (error, int)
	SalaryDataFind(w http.ResponseWriter, r *http.Request) (error, int)
	SalaryDataFindAll(w http.ResponseWriter, r *http.Request) (error, int)
}
