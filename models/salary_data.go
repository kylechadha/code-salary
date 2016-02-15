package models

import (
	"time"
)

type SalaryData struct {
	Id        int       `json:"id"`
	Company   string    `json:"company"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	Base      int       `json:"base"`
	Bonus     string    `json:"bonus"` // this can be either a percentage or a $ amount, so accept it as either for now
	Perks     int       `json:"perks"`
	Stack     []string  `json:"stack"` // maybe if you move this last?
	DateAdded time.Time `json:"dateAdded"`
}
