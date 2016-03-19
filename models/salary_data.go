package models

import (
	"time"
)

// ToDos
// - rename this all model Salary, controller Salaries... duh!

// Fields & Types
// - Remove state and country ... probably start with major cities in the US
// - perks, have it be a string of certain possible values (cell phone, gym, daycare, provided breakfast, provided lunch, snacks, etc.)
// - Bonus should also be a dollar amount

// - See the data format we get from LinkedIn, will inform what kind of data we save
// - Probably also need a user model then, certain things may be saved on the user.

// ? Switch to postgres or mongo?

// * Look into omitempty, see where else we should use it
type SalaryData struct {
	Id        int       `json:"id"`
	Company   string    `json:"company"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	Base      int       `json:"base,string,omitempty"`
	Bonus     string    `json:"bonus,omitempty"` // this can be either a percentage or a $ amount, so accept it as either for now
	Perks     int       `json:"perks"`
	Stack     []string  `json:"stack"` // maybe if you move this last?
	DateAdded time.Time `json:"dateAdded"`
}
