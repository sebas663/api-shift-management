package db

import (
	"fmt"
	"log"

	m "../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

func init() {
	session, err := mgo.Dial("127.0.0.1/api_db")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//defer session.Close()

	// optional. switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	db = session.DB("api_db")
}

func collection() *mgo.Collection {
	return db.C("apointments")
}

// GetAll returns all Apointments from the database.
func GetAll() ([]m.Apointment, error) {
	res := []m.Apointment{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// FindByID returns a single Apointment from the database.
func FindByID(id string) (*m.Apointment, error) {
	res := m.Apointment{}

	if err := collection().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Save inserts an Apointment to the database.
func Save(a m.Apointment) error {
	return collection().Insert(a)
}

// Remove deletes an Apointment from the database
func Remove(id string) error {
	//a := bson.M{"_id": id}
	//fmt.Println(a)
	err := collection().RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}
	return err
}

// Update an Apointment from the database
func Update(a m.Apointment, id string) error {
	colQuerier := bson.ObjectIdHex(id)
	change := bson.M{"$set": a}
	return collection().UpdateId(colQuerier, change)
}
