package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/micahli/notable-take-home/api/response"
	"github.com/micahli/notable-take-home/model"
)

func (a *API) getDoctorListHandler(w http.ResponseWriter, r *http.Request) {
	doctors, err := a.db.GetAllDoctors()
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if doctors == nil {
		response.Errorf(w, r, err, http.StatusBadRequest, "no doctors available")
		return
	}

	response.Write(w, r, doctors)
}

func (a *API) getAppointmentListHandler(w http.ResponseWriter, r *http.Request) {
	// Validate user input
	vars := mux.Vars(r)
	doctorUID := vars["id"]

	appointments, err := a.db.GetDoctorAppointments(doctorUID)
	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, appointments)
}

func (a *API) addAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	var ap model.Appointment
	err := json.NewDecoder(r.Body).Decode(&ap)

	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if ap.DoctorUID == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Doctor uid is missing")
		return
	} else if ap.PatientFirstName == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "patient first name is missing")
		return
	} else if ap.DoctorUID == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "patient last name is missing")
		return
	} else if ap.DateTime == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Appointment time is missing")
		return
	}

	// check the time format
	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, ap.DateTime)
	if err != nil {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Appointment time format is invalid")
		return
	}

	// check time is later than current time
	if !t.After(time.Now()) {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Appointment time is before current time")
		return
	}

	if t.Minute()%15 != 0 {
		response.Errorf(w, r, nil, http.StatusBadRequest, "Appointment time should start at 15 minutes")
		return
	}

	// try to add the appointment to db
	ap, err = a.db.AddApointment(ap)

	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Two many appointment at this time")
		return
	}

	response.Write(w, r, ap)
	return
}

func (a *API) cancelAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	var ap model.Appointment
	err := json.NewDecoder(r.Body).Decode(&ap)

	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	} else if ap.UID == "" {
		response.Errorf(w, r, nil, http.StatusBadRequest, "appointment uid is missing")
		return
	}

	// try to cancel the appointment
	err = a.db.CancelAppointment(ap)

	if err != nil {
		response.Errorf(w, r, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.Write(w, r, "ok")
	return
}
