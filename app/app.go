package app

import (
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type Ioc struct {
	ConfigService   IConfigService
	DatabaseService IDatabaseService
	DaysController  IDaysController
}

type IConfigService interface {
	GetConfig(key string) (string, error)
}

type IDatabaseService interface {
	Create(collection string, data interface{}) error
	Find(collection string, oId bson.ObjectId, document interface{}) (interface{}, error)
	FindAll(collection string) ([]interface{}, error)
}

type IDaysController interface {
	DaysCreate(w http.ResponseWriter, r *http.Request) (error, int)
	DaysFind(w http.ResponseWriter, r *http.Request) (error, int)
	DaysFindAll(w http.ResponseWriter, r *http.Request) (error, int)
}
