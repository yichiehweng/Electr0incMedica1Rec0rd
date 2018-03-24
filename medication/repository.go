package medication

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repository for type
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "EMR"

// DOCNAME the name of the document
const DOCNAME = "Medication"

//GetMedication for medication list
func (r Repository) GetMedication() Medications {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Medications{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

//AddMedication for add medication
func (r Repository) AddMedication(medication Medication) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	medication.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(medication)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//UpdateMedication for update medication
func (r Repository) UpdateMedication(medication Medication) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	medication.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(medication.ID, medication)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//DeleteMedication for Delete Medication
func (r Repository) DeleteMedication(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		return "404"
	}

	oid := bson.ObjectIdHex(id)

	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "500"
	}

	return "200"
}
