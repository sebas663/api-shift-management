package models

import "gopkg.in/mgo.v2/bson"

type (
	// Apointment represents the structure of our resource
	Apointment struct {
		ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
		PatientID   string        `json:"patientID" bson:"patientID"`
		TurnDiaryID string        `json:"turnDiaryID" bson:"turnDiaryID"`
		Coment      string        `json:"coment" bson:"coment"`
		Overturned  bool          `json:"overturned" bson:"overturned"`
		StatusCode  string        `json:"statusCode" bson:"statusCode"`
	}
	//Apointments array.
	Apointments []Apointment
)
