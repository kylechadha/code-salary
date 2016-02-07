package models

import "gopkg.in/mgo.v2/bson"

type Day struct {
	Id               bson.ObjectId `json:"id" bson:"_id"`
	DayOfTheWeek     string        `json:"dayOfTheWeek" bson:"dayOfTheWeek"`
	DayOfTheWeekType string        `json:"dayOfTheWeekType" bson:"dayOfTheWeekType"`
	SuccessfulWakeUp bool          `json:"successfulWakeUp" bson:"successfulWakeUp"`
	MorningWork      bool          `json:"morningWork" bson:"morningWork"`
	MorningWorkType  string        `json:"morningWorkType" bson:"morningWorkType"`
	WorkedOut        bool          `json:"workedOut" bson:"workedOut"`
	WorkedOutType    string        `json:"workedOutType" bson:"workedOutType"`
	PlannedNextDay   bool          `json:"plannedNextDay" bson:"plannedNextDay"`
}
