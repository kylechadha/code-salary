package models

type SalaryData struct {
	// Id               bson.ObjectId `json:"id" bson:"_id"`
	Company  string   `json:"company"`
	Location string   `json:"location"`
	Base     int      `json:"base"`
	Bonus    string   `json:"bonus"` // this can be either a percentage or a $ amount, so accept it as either for now
	Perks    int      `json:"perks"`
	Stack    []string `json:"stack"`
}
