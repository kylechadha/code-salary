package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SalaryData struct {
	Id        bson.ObjectId `json:"id"`
	Company   string        `json:"company"`
	City      string        `json:"city"`
	State     string        `json:"state"`
	Country   string        `json:"country"`
	Base      int           `json:"base"`
	Bonus     string        `json:"bonus"` // this can be either a percentage or a $ amount, so accept it as either for now
	Perks     int           `json:"perks"`
	Stack     []string      `json:"stack"`
	DateAdded time.Time     `json:"dateAdded"`
}
