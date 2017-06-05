package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	db "../db"
	m "../models"
	"github.com/gorilla/mux"
)

//ApointmentIndex root of the api.
func ApointmentIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	apointmentArray, err := db.GetAll()
	if err = json.NewEncoder(w).Encode(apointmentArray); err != nil {
		panic(err)
	}
}

//ApointmentSave create an Apointment.
func ApointmentSave(w http.ResponseWriter, r *http.Request) {
	var apointment m.Apointment
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &apointment); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := db.Save(apointment)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

//ApointmentFindByID find by id.
func ApointmentFindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apointmentID := vars["apointmentID"]
	var err error
	if len(apointmentID) == 0 {
		panic(err)
	}
	apointment, err := db.FindByID(apointmentID)
	if len(apointment.ID) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(apointment); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(m.JSONErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

//ApointmentUpdate find by id.
func ApointmentUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apointmentID := vars["apointmentID"]
	var apointment m.Apointment
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &apointment); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := db.Update(apointment, apointmentID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return
}

//ApointmentDelete find by id.
func ApointmentDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apointmentID := vars["apointmentID"]

	if len(apointmentID) == 0 {
		var err error
		panic(err)
	}

	err := db.Remove(apointmentID)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(m.JSONErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
