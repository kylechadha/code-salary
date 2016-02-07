package services

import (
	"fmt"

	"github.com/kylechadha/omnia-app/app"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Database Service type.
type databaseService struct {
	session *mgo.Session
}

// Database Service constructor function.
func NewDatabaseService(app *app.Ioc) *databaseService {
	mongoUrl, err := app.ConfigService.GetConfig("mongourl")
	if err != nil {
		fmt.Println("Config file does not include a 'mongourl'.")
		panic(err)
	}

	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println("Unable to connect to the Mongo DB.")
		panic(err)
	}

	return &databaseService{session}
}

func (d *databaseService) Create(collection string, document interface{}) error {

	// Write the document to the database.
	err := d.session.DB("omnia").C(collection).Insert(document)
	if err != nil {
		return err
	}

	return nil
}

func (d *databaseService) Find(collection string, oId bson.ObjectId, document interface{}) (interface{}, error) {

	// Find the document by ID.
	err := d.session.DB("omnia").C(collection).FindId(oId).One(&document)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (d *databaseService) FindAll(collection string) ([]interface{}, error) {

	// Find all documents in the collection.
	var documents []interface{}
	err := d.session.DB("omnia").C(collection).Find(bson.M{}).All(&documents)
	if err != nil {
		return nil, err
	}

	return documents, nil
}
