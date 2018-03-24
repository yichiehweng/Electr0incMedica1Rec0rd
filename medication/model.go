package medication

import "gopkg.in/mgo.v2/bson"

//Medication from MongoDB
type Medication struct {
	ID        bson.ObjectId `bson:"_id"`
	ATCCode   string        `bson:"ATCCode"`
	Name      string        `bson:"Name"`
	TradeName string        `bson:"TradeName"`
	Category  string        `bson:"Category"`
}

//Medications for list of medications
type Medications []Medication
